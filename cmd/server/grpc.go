package main

import (
	//"fmt"
	//"github.com/Infoblox-CTO/go.grpc.middleware/authz"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/seizadi/cmdb/pkg/pb"
	"github.com/seizadi/cmdb/pkg/svc"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	//"github.com/Infoblox-CTO/go.grpc.middleware/authz"
	"github.com/infobloxopen/atlas-app-toolkit/auth"
	"github.com/infobloxopen/atlas-app-toolkit/errors"
	"github.com/infobloxopen/atlas-app-toolkit/errors/mappers/validationerrors"
	"github.com/infobloxopen/atlas-app-toolkit/requestid"
)

func NewGRPCServer(logger *logrus.Logger, db *gorm.DB) (*grpc.Server, error) {
	interceptors := []grpc.UnaryServerInterceptor{
		// logging middleware
		grpc_logrus.UnaryServerInterceptor(logrus.NewEntry(logger)),
		
		// Request-Id interceptor
		requestid.UnaryServerInterceptor(),
		auth.LogrusUnaryServerInterceptor(),
		errors.UnaryServerInterceptor(ErrorMappings...),
		// validation middleware
		validationerrors.UnaryServerInterceptor(),
		//UnaryServerInterceptor(),
	}
	// add authorization interceptor if authz service address is provided
	//if viper.GetBool("atlas.authz.enable") {
	//	// authorization interceptor
	//	interceptors = append(interceptors, authz.UnaryServerInterceptor(
	//		fmt.Sprintf("%s:%s", viper.GetString("atlas.authz.address"), viper.GetString("atlas.authz.port")), "cmdb"),
	//	)
	//}

	// create new gRPC grpcServer with middleware chain
	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(interceptors...)))

	// register all of our services with the grpcServer
	s, err := svc.NewBasicServer(db)
	if err != nil {
		return nil, err
	}
	pb.RegisterCmdbServer(grpcServer, s)

	return grpcServer, nil
}

//func UnaryServerInterceptor() grpc.UnaryServerInterceptor {
//
//	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
//		f := &query.Filtering{}
//		err := gateway.GetCollectionOp(req, f)
//		if err != nil {
//			return nil, err
//		}
//		s := &query.Sorting{}
//		err = gateway.GetCollectionOp(req, s)
//		if err != nil {
//			return nil, err
//		}
//		fs := &query.FieldSelection{}
//		err = gateway.GetCollectionOp(req, fs)
//		if err != nil {
//			return nil, err
//		}
//		return handler(ctx, req)
//	}
//}
