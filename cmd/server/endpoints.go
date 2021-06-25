package main

import (
	"github.com/infobloxopen/atlas-app-toolkit/gateway"
	
	"github.com/seizadi/cmdb/pkg/pb"
)

func RegisterGatewayEndpoints(prefix string) gateway.Option {
	return gateway.WithEndpointRegistration(prefix,
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
		pb.RegisterNetworksHandlerFromEndpoint,
		pb.RegisterManifestHandlerFromEndpoint,
	)
}
