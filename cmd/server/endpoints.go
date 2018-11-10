package main

import (
	"github.com/infobloxopen/atlas-app-toolkit/gateway"
	"github.com/spf13/viper"
	
	"github.com/seizadi/cmdb/pkg/pb"
)

func RegisterGatewayEndpoints() gateway.Option {
	return gateway.WithEndpointRegistration(viper.GetString("server.version"),
		pb.RegisterCmdbHandlerFromEndpoint,
		pb.RegisterAwsServicesHandlerFromEndpoint,
		pb.RegisterContainersHandlerFromEndpoint,
		pb.RegisterManifestsHandlerFromEndpoint,
		pb.RegisterVersionTagsHandlerFromEndpoint,
		pb.RegisterArtifactsHandlerFromEndpoint,
		pb.RegisterAwsRdsInstancesHandlerFromEndpoint,
		pb.RegisterDeploymentsHandlerFromEndpoint,
		pb.RegisterEnvironmentsHandlerFromEndpoint,
		pb.RegisterKubeClustersHandlerFromEndpoint,
		pb.RegisterRegionsHandlerFromEndpoint,
		pb.RegisterSecretsHandlerFromEndpoint,
		pb.RegisterVaultsHandlerFromEndpoint,
		pb.RegisterApplicationsHandlerFromEndpoint,
	)
}
