package main

import (
	"flag"
	"net"

	"github.com/infobloxopen/atlas-app-toolkit/gateway"
	"github.com/infobloxopen/atlas-app-toolkit/server"
	"github.com/seizadi/cmdb/cmd"
	"github.com/seizadi/cmdb/pkg/pb"
	"github.com/sirupsen/logrus"
)

var (
	ServerAddress      string
	GatewayAddress     string
	SwaggerDir         string
	DBConnectionString string
)

func main() {
	logger := logrus.New()
	grpcServer, err := NewGRPCServer(logger, DBConnectionString)
	if err != nil {
		logger.Fatalln(err)
	}

	s, err := server.NewServer(
		server.WithGrpcServer(grpcServer),
		server.WithGateway(
			gateway.WithServerAddress(ServerAddress),
			gateway.WithEndpointRegistration(cmd.GatewayURL, pb.RegisterCmdbHandlerFromEndpoint),
		),
		server.WithHandler("/swagger/", NewSwaggerHandler(SwaggerDir)),
	)
	if err != nil {
		logger.Fatalln(err)
	}

	grpcL, err := net.Listen("tcp", ServerAddress)
	if err != nil {
		logger.Fatalln(err)
	}

	httpL, err := net.Listen("tcp", GatewayAddress)
	if err != nil {
		logger.Fatalln(err)
	}

	logger.Printf("serving gRPC at %s", ServerAddress)
	logger.Printf("serving http at %s", GatewayAddress)

	if err := s.Serve(grpcL, httpL); err != nil {
		logger.Fatalln(err)
	}
}

func init() {
	// default server address; optionally set via command-line flags
	flag.StringVar(&ServerAddress, "address", cmd.ServerAddress, "the gRPC server address")
	flag.StringVar(&GatewayAddress, "gateway", cmd.GatewayAddress, "address of the gateway server")
	flag.StringVar(&SwaggerDir, "swagger-dir", "pkg/pb/service.swagger.json", "directory of the swagger.json file")
	flag.StringVar(&DBConnectionString, "db", cmd.DBConnectionString, "the database address")
	flag.Parse()
}
