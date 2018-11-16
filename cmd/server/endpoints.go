package main

import (
	"github.com/infobloxopen/atlas-app-toolkit/gateway"
	"github.com/spf13/viper"
	
	"github.com/seizadi/cmdb/pkg/pb"
)

func RegisterGatewayEndpoints() gateway.Option {
	return gateway.WithEndpointRegistration(viper.GetString("server.version"),
		pb.RegisterCmdbHandlerFromEndpoint,
		pb.RegisterApplicationsHandlerFromEndpoint,
		pb.RegisterAwsServicesHandlerFromEndpoint,
		pb.RegisterRegionsHandlerFromEndpoint,
		pb.RegisterVaultsHandlerFromEndpoint,
		pb.RegisterArtifactsHandlerFromEndpoint,
		pb.RegisterSecretsHandlerFromEndpoint,
		pb.RegisterAwsRdsInstancesHandlerFromEndpoint,
		pb.RegisterDeploymentsHandlerFromEndpoint,
		pb.RegisterKubeClustersHandlerFromEndpoint,
		pb.RegisterManifestsHandlerFromEndpoint,
		pb.RegisterVersionTagsHandlerFromEndpoint,
		pb.RegisterContainersHandlerFromEndpoint,
		pb.RegisterEnvironmentsHandlerFromEndpoint,
		pb.RegisterCloudProvidersHandlerFromEndpoint,
	)
}
