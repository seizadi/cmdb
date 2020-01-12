package main

import (
	"github.com/infobloxopen/atlas-app-toolkit/gateway"
	"github.com/spf13/viper"

	"github.com/seizadi/cmdb/pkg/pb"
)

func RegisterGatewayEndpoints() gateway.Option {
	return gateway.WithEndpointRegistration(viper.GetString("server.version"),
		pb.RegisterCmdbHandlerFromEndpoint,
		pb.RegisterCloudProvidersHandlerFromEndpoint,
		pb.RegisterDeploymentsHandlerFromEndpoint,
		pb.RegisterRegionsHandlerFromEndpoint,
		pb.RegisterEnvironmentsHandlerFromEndpoint,
		pb.RegisterApplicationsHandlerFromEndpoint,
		pb.RegisterApplicationInstancesHandlerFromEndpoint,
		pb.RegisterAppVersionsHandlerFromEndpoint,
		pb.RegisterLifecyclesHandlerFromEndpoint,
		pb.RegisterAppConfigsHandlerFromEndpoint,
		pb.RegisterArtifactsHandlerFromEndpoint,
		pb.RegisterKubeClustersHandlerFromEndpoint,
		pb.RegisterSecretsHandlerFromEndpoint,
		pb.RegisterVaultsHandlerFromEndpoint,
		pb.RegisterChartVersionsHandlerFromEndpoint,
	)
}
