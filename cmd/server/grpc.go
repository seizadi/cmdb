package main

import (
	"context"
	"io/ioutil"
	
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus/ctxlogrus"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	
	"github.com/infobloxopen/atlas-app-toolkit/auth"
	"github.com/infobloxopen/atlas-app-toolkit/errors"
	"github.com/infobloxopen/atlas-app-toolkit/errors/mappers/validationerrors"
	"github.com/infobloxopen/atlas-app-toolkit/gateway"
	tkgorm "github.com/infobloxopen/atlas-app-toolkit/gorm"
	"github.com/infobloxopen/atlas-app-toolkit/logging"
	"github.com/infobloxopen/atlas-app-toolkit/query"
	"github.com/infobloxopen/atlas-app-toolkit/requestid"
)

func NewGRPCServer(logger *logrus.Logger, db *gorm.DB) (*grpc.Server, error) {
	interceptors := []grpc.UnaryServerInterceptor{
		grpc_logrus.UnaryServerInterceptor(logrus.NewEntry(logger)),
		requestid.UnaryServerInterceptor(),
		dbLoggingWrapper(db),
		auth.LogrusUnaryServerInterceptor(),
		errors.UnaryServerInterceptor(ErrorMappings...),
		// validation interceptor
		validationerrors.UnaryServerInterceptor(),
		UnaryServerInterceptor(),
	}

	// if per request log level is included, insert it right after the grpc_logrus interceptor
	if viper.GetBool("logging.dynamic.level") {
		interceptors = append(interceptors, nil)
		copy(interceptors[2:], interceptors[1:])
		interceptors[1] = logging.LogLevelInterceptor(logger.Level)
	}
	
	return CreateServer(logger, db, interceptors)

}

func UnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		f := &query.Filtering{}
		err := gateway.GetCollectionOp(req, f)
		if err != nil {
			return nil, err
		}
		s := &query.Sorting{}
		err = gateway.GetCollectionOp(req, s)
		if err != nil {
			return nil, err
		}
		fs := &query.FieldSelection{}
		err = gateway.GetCollectionOp(req, fs)
		if err != nil {
			return nil, err
		}

		return handler(ctx, req)
	}
}

// creates a per-request copy of the DB with db logging set using the context logger
func dbLoggingWrapper(db *gorm.DB) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		logEntry := ctxlogrus.Extract(ctx)
		// Do nothing if no logger was found in the context
		if logEntry.Logger.Out != ioutil.Discard && viper.GetBool("database.logging") {
			db = db.New()
			db.SetLogger(logEntry)
			db.LogMode(true)
		}
		return tkgorm.UnaryServerInterceptor(db)(ctx, req, info, handler)
	}
}
