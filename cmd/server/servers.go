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

	application, err := svc.NewApplicationsServer()
	if err != nil {
		return nil, err
	}
	pb.RegisterApplicationsServer(grpcServer, application)

	aws_service, err := svc.NewAwsServicesServer()
	if err != nil {
		return nil, err
	}
	pb.RegisterAwsServicesServer(grpcServer, aws_service)

	region, err := svc.NewRegionsServer()
	if err != nil {
		return nil, err
	}
	pb.RegisterRegionsServer(grpcServer, region)

	vault, err := svc.NewVaultsServer()
	if err != nil {
		return nil, err
	}
	pb.RegisterVaultsServer(grpcServer, vault)

	artifact, err := svc.NewArtifactsServer()
	if err != nil {
		return nil, err
	}
	pb.RegisterArtifactsServer(grpcServer, artifact)

	secret, err := svc.NewSecretsServer()
	if err != nil {
		return nil, err
	}
	pb.RegisterSecretsServer(grpcServer, secret)

	aws_rds_instance, err := svc.NewAwsRdsInstancesServer()
	if err != nil {
		return nil, err
	}
	pb.RegisterAwsRdsInstancesServer(grpcServer, aws_rds_instance)

	deployment, err := svc.NewDeploymentsServer()
	if err != nil {
		return nil, err
	}
	pb.RegisterDeploymentsServer(grpcServer, deployment)

	kube_cluster, err := svc.NewKubeClustersServer()
	if err != nil {
		return nil, err
	}
	pb.RegisterKubeClustersServer(grpcServer, kube_cluster)

	manifest, err := svc.NewManifestsServer()
	if err != nil {
		return nil, err
	}
	pb.RegisterManifestsServer(grpcServer, manifest)

	version_tag, err := svc.NewVersionTagsServer()
	if err != nil {
		return nil, err
	}
	pb.RegisterVersionTagsServer(grpcServer, version_tag)

	container, err := svc.NewContainersServer()
	if err != nil {
		return nil, err
	}
	pb.RegisterContainersServer(grpcServer, container)

	environment, err := svc.NewEnvironmentsServer()
	if err != nil {
		return nil, err
	}
	pb.RegisterEnvironmentsServer(grpcServer, environment)

	cloud_provider, err := svc.NewCloudProvidersServer()
	if err != nil {
		return nil, err
	}
	pb.RegisterCloudProvidersServer(grpcServer, cloud_provider)

	return grpcServer, nil
}
