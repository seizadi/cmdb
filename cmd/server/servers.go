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

	application, err := svc.NewApplicationsServer()
	if err != nil {
		return nil, err
	}
	pb.RegisterApplicationsServer(grpcServer, application)

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

	kube_cluster, err := svc.NewKubeClustersServer()
	if err != nil {
		return nil, err
	}
	pb.RegisterKubeClustersServer(grpcServer, kube_cluster)

	app_version, err := svc.NewAppVersionsServer()
	if err != nil {
		return nil, err
	}
	pb.RegisterAppVersionsServer(grpcServer, app_version)

	lifecycle, err := svc.NewLifecyclesServer()
	if err != nil {
		return nil, err
	}
	pb.RegisterLifecyclesServer(grpcServer, lifecycle)

	app_config, err := svc.NewAppConfigsServer()
	if err != nil {
		return nil, err
	}
	pb.RegisterAppConfigsServer(grpcServer, app_config)

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

	chart_version, err := svc.NewChartVersionsServer()
	if err != nil {
		return nil, err
	}
	pb.RegisterChartVersionsServer(grpcServer, chart_version)

	return grpcServer, nil
}
