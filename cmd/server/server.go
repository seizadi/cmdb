package main

import (
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/jinzhu/gorm"
	"github.com/seizadi/cmdb/pkg/pb"
	"github.com/seizadi/cmdb/pkg/svc"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func CreateServer(logger *logrus.Logger, db *gorm.DB, interceptors []grpc.UnaryServerInterceptor) (*grpc.Server, error) {
	// create new gRPC grpcServer with middleware chain
	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(interceptors...)))
	
	// register all of our services into the grpcServer
	s, err := svc.NewBasicServer(db)
	if err != nil {
		return nil, err
	}
	pb.RegisterCmdbServer(grpcServer, s)
	
	rs, err := svc.NewRegionsServer()
	if err != nil {
		return nil, err
	}
	pb.RegisterRegionsServer(grpcServer, rs)
	
	cs, err := svc.NewContainersServer()
	if err != nil {
		return nil, err
	}
	pb.RegisterContainersServer(grpcServer, cs)
	
	vts, err := svc.NewVersionTagsServer()
	if err != nil {
		return nil, err
	}
	pb.RegisterVersionTagsServer(grpcServer, vts)
	
	ss, err := svc.NewSecretsServer()
	if err != nil {
		return nil, err
	}
	pb.RegisterSecretsServer(grpcServer, ss)
	
	vs, err := svc.NewVaultsServer()
	if err != nil {
		return nil, err
	}
	pb.RegisterVaultsServer(grpcServer, vs)
	
	return grpcServer, nil
}
