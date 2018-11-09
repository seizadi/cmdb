package main

import (
	"github.com/infobloxopen/atlas-app-toolkit/gateway"
	"github.com/spf13/viper"
	
	"github.com/seizadi/cmdb/pkg/pb"
)

func RegisterGatewayEndpoints() gateway.Option {
	return gateway.WithEndpointRegistration(viper.GetString("server.version"),
		pb.RegisterCmdbHandlerFromEndpoint,
		pb.RegisterRegionsHandlerFromEndpoint,
		pb.RegisterContainersHandlerFromEndpoint,
		pb.RegisterVersionTagsHandlerFromEndpoint,
		pb.RegisterSecretsHandlerFromEndpoint,
		pb.RegisterVaultsHandlerFromEndpoint,
	)
}
