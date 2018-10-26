package main

import (
	"context"
	//"fmt"
	//"github.com/Infoblox-CTO/go.grpc.middleware/authz"
	
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/seizadi/cmdb/pkg/pb"
	"github.com/seizadi/cmdb/pkg/svc"
	//"github.com/Infoblox-CTO/go.grpc.middleware/authz"
	"github.com/infobloxopen/atlas-app-toolkit/auth"
	"github.com/infobloxopen/atlas-app-toolkit/errors"
	"github.com/infobloxopen/atlas-app-toolkit/errors/mappers/validationerrors"
	"github.com/infobloxopen/atlas-app-toolkit/gateway"
	"github.com/infobloxopen/atlas-app-toolkit/query"
	"github.com/infobloxopen/atlas-app-toolkit/requestid"
)

func NewGRPCServer(logger *logrus.Logger, db *gorm.DB) (*grpc.Server, error) {
	interceptors := []grpc.UnaryServerInterceptor{
		grpc_logrus.UnaryServerInterceptor(logrus.NewEntry(logger)),
		requestid.UnaryServerInterceptor(),
		auth.LogrusUnaryServerInterceptor(),
		errors.UnaryServerInterceptor(ErrorMappings...),
		// validation interceptor
		validationerrors.UnaryServerInterceptor(),
		UnaryServerInterceptor(),
	}
	// add authorization interceptor if authz service address is provided
	//if viper.GetBool("atlas.authz.enable") {
	//	// authorization interceptor
	//	interceptors = append(interceptors, authz.UnaryServerInterceptor(
	//		fmt.Sprintf("%s:%s", viper.GetString("atlas.authz.address"), viper.GetString("atlas.authz.port")), "contacts"),
	//	)
	//}

	// create new gRPC grpcServer with middleware chain
	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(interceptors...)))

	// register all of our services into the grpcServer
	ps, err := svc.NewProfilesServer(db)
	if err != nil {
		return nil, err
	}
	pb.RegisterProfilesServer(grpcServer, ps)

	gs, err := svc.NewGroupsServer(db)
	if err != nil {
		return nil, err
	}
	pb.RegisterGroupsServer(grpcServer, gs)

	cs, err := svc.NewContactsServer(db)
	if err != nil {
		return nil, err
	}
	pb.RegisterContactsServer(grpcServer, cs)

	return grpcServer, nil
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
		if err := pb.ContactsValidateFiltering(info.FullMethod, f); err != nil {
			return nil, status.Errorf(codes.InvalidArgument, err.Error())
		}
		if err := pb.ContactsValidateSorting(info.FullMethod, s); err != nil {
			return nil, status.Errorf(codes.InvalidArgument, err.Error())
		}
		if err := pb.ContactsValidateFieldSelection(info.FullMethod, fs); err != nil {
			return nil, status.Errorf(codes.InvalidArgument, err.Error())
		}
		return handler(ctx, req)
	}
}
