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
	
	cloud_provider, err := svc.NewCloudProvidersServer()
	if err != nil {
		return nil, err
	}
	pb.RegisterCloudProvidersServer(grpcServer, cloud_provider)
	
	deployment, err := svc.NewDeploymentsServer()
	if err != nil {
		return nil, err
	}
	pb.RegisterDeploymentsServer(grpcServer, deployment)
	
	region, err := svc.NewRegionsServer()
	if err != nil {
		return nil, err
	}
	pb.RegisterRegionsServer(grpcServer, region)
	
	stage, err := svc.NewStagesServer()
	if err != nil {
		return nil, err
	}
	pb.RegisterStagesServer(grpcServer, stage)
	
	environment, err := svc.NewEnvironmentsServer()
	if err != nil {
		return nil, err
	}
	pb.RegisterEnvironmentsServer(grpcServer, environment)
	
	application_instance, err := svc.NewApplicationInstancesServer()
	if err != nil {
		return nil, err
	}
	pb.RegisterApplicationInstancesServer(grpcServer, application_instance)
	
	artifact, err := svc.NewArtifactsServer()
	if err != nil {
		return nil, err
	}
	pb.RegisterArtifactsServer(grpcServer, artifact)
	
	application, err := svc.NewApplicationsServer()
	if err != nil {
		return nil, err
	}
	pb.RegisterApplicationsServer(grpcServer, application)
	
	kube_cluster, err := svc.NewKubeClustersServer()
	if err != nil {
		return nil, err
	}
	pb.RegisterKubeClustersServer(grpcServer, kube_cluster)
	
	aws_rds_instance, err := svc.NewAwsRdsInstancesServer()
	if err != nil {
		return nil, err
	}
	pb.RegisterAwsRdsInstancesServer(grpcServer, aws_rds_instance)
	
	secret, err := svc.NewSecretsServer()
	if err != nil {
		return nil, err
	}
	pb.RegisterSecretsServer(grpcServer, secret)
	
	vault, err := svc.NewVaultsServer()
	if err != nil {
		return nil, err
	}
	pb.RegisterVaultsServer(grpcServer, vault)
	
	value, err := svc.NewValuesServer()
	if err != nil {
		return nil, err
	}
	pb.RegisterValuesServer(grpcServer, value)
	
	aws_service, err := svc.NewAwsServicesServer()
	if err != nil {
		return nil, err
	}
	pb.RegisterAwsServicesServer(grpcServer, aws_service)
	
	chart_version, err := svc.NewChartVersionsServer()
	if err != nil {
		return nil, err
	}
	pb.RegisterChartVersionsServer(grpcServer, chart_version)
	
	return grpcServer, nil
}
