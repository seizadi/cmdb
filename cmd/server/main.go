package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
	"sync"
	"text/template"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"google.golang.org/grpc"

	"github.com/infobloxopen/atlas-app-toolkit/gateway"
	"github.com/infobloxopen/atlas-app-toolkit/gorm/resource"
	"github.com/infobloxopen/atlas-app-toolkit/health"
	"github.com/infobloxopen/atlas-app-toolkit/server"
	atlas_validate "github.com/infobloxopen/protoc-gen-atlas-validate/interceptor"

	"github.com/seizadi/cmdb/pkg/pb"
)

var (
	SwaggerPatch sync.Once
)

func init() {
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)
	viper.AutomaticEnv()
	viper.SetEnvPrefix("CMDB")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AddConfigPath(viper.GetString("config.source"))
	if viper.GetString("config.file") != "" {
		viper.SetConfigName(viper.GetString("config.file"))
		if err := viper.ReadInConfig(); err != nil {
			log.Fatalf("cannot load configuration: %v", err)
		}
	}
	resource.RegisterApplication(viper.GetString("app.id"))
	resource.SetPlural()
	if viper.GetString("database.dsn") == "" {
		setDBString()
	}
}

func main() {
	doneC := make(chan error)
	logger := NewLogger()

	go func() { doneC <- ServeInternal(logger) }()
	go func() { doneC <- ServeExternal(logger) }()

	if err := <-doneC; err != nil {
		logger.Fatal(err)
	}
}

func NewLogger() *logrus.Logger {
	logger := logrus.StandardLogger()

	// Set the log level on the default logger based on command line flag
	logLevels := map[string]logrus.Level{
		"debug":   logrus.DebugLevel,
		"info":    logrus.InfoLevel,
		"warning": logrus.WarnLevel,
		"error":   logrus.ErrorLevel,
		"fatal":   logrus.FatalLevel,
		"panic":   logrus.PanicLevel,
	}
	if level, ok := logLevels[viper.GetString("logging.level")]; !ok {
		logger.Errorf("Invalid %q provided for log level", viper.GetString("logging.level"))
		logger.SetLevel(logrus.InfoLevel)
	} else {
		logger.SetLevel(level)
	}

	return logger
}

// ServeInternal builds and runs the server that listens on InternalAddress
func ServeInternal(logger *logrus.Logger) error {
	healthChecker := health.NewChecksHandler(
		viper.GetString("internal.health"),
		viper.GetString("internal.readiness"),
	)
	healthChecker.AddReadiness("DB ready check", dbReady)
	healthChecker.AddLiveness("ping", health.HTTPGetCheck(
		fmt.Sprint("http://", viper.GetString("internal.address"), ":", viper.GetString("internal.port"), "/ping"), time.Minute),
	)

	s, err := server.NewServer(
		// register our health checks
		server.WithHealthChecks(healthChecker),
		// this endpoint will be used for our health checks
		server.WithHandler("/ping", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(200)
			w.Write([]byte("pong"))
		})),
	)
	if err != nil {
		return err
	}
	l, err := net.Listen("tcp", fmt.Sprintf("%s:%s", viper.GetString("internal.address"), viper.GetString("internal.port")))
	if err != nil {
		return err
	}

	logger.Debugf("serving internal http at %q", fmt.Sprintf("%s:%s", viper.GetString("internal.address"), viper.GetString("internal.port")))

	return s.Serve(nil, l)
}

func forwardResponseOption(ctx context.Context, w http.ResponseWriter, resp proto.Message) error {
	w.Header().Set("Cache-Control", "no-cache, no-store, max-age=0, must-revalidate")
	return nil
}

// ServeExternal builds and runs the server that listens on ServerAddress and GatewayAddress
func ServeExternal(logger *logrus.Logger) error {

	dbSQL, err := sql.Open(viper.GetString("database.type"), viper.GetString("database.dsn"))
	if err != nil {
		return err
	}
	defer dbSQL.Close()
	db, err := gorm.Open("postgres", dbSQL)
	if err != nil {
		return err
	}
	defer db.Close()

	grpcServer, err := NewGRPCServer(logger, db)
	if err != nil {
		return err
	}
	
	prefix := "/" + viper.GetString("gateway.endpoint") + viper.GetString("server.version")
	prefixApiDoc := "/" + viper.GetString("gateway.endpoint") + "/apidoc"
	
	s, err := server.NewServer(
		// register our grpc server
		server.WithGrpcServer(grpcServer),
		// register the gateway to proxy to the given server address with the service registration endpoints
		server.WithGateway(
			gateway.WithGatewayOptions(
				runtime.WithMetadata(pb.AtlasValidateAnnotator),
				runtime.WithMetadata(gateway.NewPresenceAnnotator("PATCH")),
				runtime.WithForwardResponseOption(forwardResponseOption),
			),
			gateway.WithDialOptions(
				[]grpc.DialOption{grpc.WithInsecure(), grpc.WithUnaryInterceptor(
					grpc_middleware.ChainUnaryClient(
						[]grpc.UnaryClientInterceptor{
							atlas_validate.ValidationClientInterceptor(),
							gateway.ClientUnaryInterceptor,
							gateway.PresenceClientInterceptor()}...,
					),
				)}...,
			),
			gateway.WithServerAddress(fmt.Sprintf("%s:%s", viper.GetString("server.address"), viper.GetString("server.port"))),
			RegisterGatewayEndpoints(prefix),
		),
		// serve swagger at the root
		server.WithHandler("/swagger", http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			basePath := prefix + "/"
			if origURI := request.Header.Get("X-Original-Uri"); origURI != "" {
				// k8s deployment patch.
				SwaggerPatch.Do(func() {
					p := strings.TrimSuffix(origURI, "/swagger")
					basePath = p + basePath

					var v map[string]interface{}
					b, err := ioutil.ReadFile(viper.GetString("gateway.swaggerFile"))
					if err != nil {
						http.Error(writer, "Internal Error", 500)
						return
					}

					if err = json.Unmarshal(b, &v); err != nil {
						http.Error(writer, "Internal Error", 500)
						return
					}

					v["basePath"] = basePath

					if v, ok := v["securityDefinitions"].(map[string]interface{}); ok {
						if v, ok := v["ApiKeyAuth"].(map[string]interface{}); ok {
							v["name"] = "User-And-Pass"
						}
					}

					v["schemes"] = []string{"https"}

					f, err := os.Create(viper.GetString("gateway.swaggerFile") + ".patched")
					if err != nil {
						logger.Debugf("Unable to create %q: %v", viper.GetString("gateway.swaggerFile")+".patched", err)
						http.Error(writer, "Internal Error", 500)
						return
					}

					if err = json.NewEncoder(f).Encode(v); err != nil {
						logger.Debugf("Unable to encode json: %v", err)
						http.Error(writer, "Internal Error", 500)
						return
					}

					f.Close()
				})

				http.ServeFile(writer, request, viper.GetString("gateway.swaggerFile")+".patched")

			} else {
				http.ServeFile(writer, request, viper.GetString("gateway.swaggerFile"))
				return
			}
		})),

		server.WithHandler(prefixApiDoc + "/", http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			filenames := strings.Split(request.URL.Path, "/")
			http.ServeFile(writer, request, viper.GetString("gateway.swaggerUI")+filenames[len(filenames)-1])
		})),

		server.WithHandler(prefixApiDoc, http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			host := strings.TrimSuffix(request.Host, ".")

			var schema string

			// We got to figure out the original URI and schema client was requesting
			// to insert it into the index.html for css/html sources.
			if origURI := request.Header.Get("X-Original-Uri"); origURI != "" {
				host += strings.TrimSuffix(origURI, prefixApiDoc)
			} else {
				host += strings.TrimSuffix(request.URL.Path, prefixApiDoc)
			}

			if origSchema := request.Header.Get("X-Scheme"); origSchema != "" {
				schema = origSchema + "://"
			} else {
				if request.URL.Scheme == "http" || request.URL.Scheme == "" {
					schema = "http://"
				} else {
					schema = "https://"
				}
			}

			t, err := template.ParseFiles(viper.GetString("gateway.swaggerUI") + "index.html.tt")
			if err != nil {
				logger.Debugf("Error wile rendering template: %v", err)
				return
			}
			t.Execute(writer, struct{ HRef, SchemaRef string }{HRef: schema + host + prefixApiDoc, SchemaRef: schema + host + "/swagger"})
		})),
	)
	if err != nil {
		return err
	}

	// open some listeners for our server and gateway
	grpcL, err := net.Listen("tcp", fmt.Sprintf("%s:%s", viper.GetString("server.address"), viper.GetString("server.port")))
	if err != nil {
		return err
	}
	gatewayL, err := net.Listen("tcp", fmt.Sprintf("%s:%s", viper.GetString("gateway.address"), viper.GetString("gateway.port")))
	if err != nil {
		return err
	}

	logger.Debugf("serving gRPC at %q", fmt.Sprintf("%s:%s", viper.GetString("server.address"), viper.GetString("server.port")))
	logger.Debugf("serving http at %q", fmt.Sprintf("%s:%s", viper.GetString("gateway.address"), viper.GetString("gateway.port")))

	return s.Serve(grpcL, gatewayL)
}

func setDBString() {
	viper.Set("database.dsn", fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		viper.GetString("database.user"), viper.GetString("database.password"),
		viper.GetString("database.host"), viper.GetString("database.port"),
		viper.GetString("database.name"), viper.GetString("database.ssl"),
	))
}

func dbReady() error {
	db, err := sql.Open(viper.GetString("database.type"), viper.GetString("database.dsn"))
	if err != nil {
		return err
	}
	defer db.Close()
	return db.Ping()
}
