package svc

import (
	"github.com/seizadi/cmdb/pkg/pb"
)

type cloudProvidersServer struct {
	*pb.CloudProvidersDefaultServer
}

// NewCloudProvidersServer returns an instance of the default CloudProvider server interface
func NewCloudProvidersServer() (pb.CloudProvidersServer, error) {
	return &cloudProvidersServer{&pb.CloudProvidersDefaultServer{}}, nil
}

type deploymentsServer struct {
	*pb.DeploymentsDefaultServer
}

// NewDeploymentsServer returns an instance of the default Deployment server interface
func NewDeploymentsServer() (pb.DeploymentsServer, error) {
	return &deploymentsServer{&pb.DeploymentsDefaultServer{}}, nil
}

type regionsServer struct {
	*pb.RegionsDefaultServer
}

// NewRegionsServer returns an instance of the default Region server interface
func NewRegionsServer() (pb.RegionsServer, error) {
	return &regionsServer{&pb.RegionsDefaultServer{}}, nil
}

type environmentsServer struct {
	*pb.EnvironmentsDefaultServer
}

// NewEnvironmentsServer returns an instance of the default Environment server interface
func NewEnvironmentsServer() (pb.EnvironmentsServer, error) {
	return &environmentsServer{&pb.EnvironmentsDefaultServer{}}, nil
}

type applicationInstancesServer struct {
	*pb.ApplicationInstancesDefaultServer
}

// NewApplicationInstancesServer returns an instance of the default ApplicationInstance server interface
func NewApplicationInstancesServer() (pb.ApplicationInstancesServer, error) {
	return &applicationInstancesServer{&pb.ApplicationInstancesDefaultServer{}}, nil
}

type artifactsServer struct {
	*pb.ArtifactsDefaultServer
}

// NewArtifactsServer returns an instance of the default Artifact server interface
func NewArtifactsServer() (pb.ArtifactsServer, error) {
	return &artifactsServer{&pb.ArtifactsDefaultServer{}}, nil
}

type kubeClustersServer struct {
	*pb.KubeClustersDefaultServer
}

// NewKubeClustersServer returns an instance of the default KubeCluster server interface
func NewKubeClustersServer() (pb.KubeClustersServer, error) {
	return &kubeClustersServer{&pb.KubeClustersDefaultServer{}}, nil
}

type secretsServer struct {
	*pb.SecretsDefaultServer
}

// NewSecretsServer returns an instance of the default Secret server interface
func NewSecretsServer() (pb.SecretsServer, error) {
	return &secretsServer{&pb.SecretsDefaultServer{}}, nil
}

type vaultsServer struct {
	*pb.VaultsDefaultServer
}

// NewVaultsServer returns an instance of the default Vault server interface
func NewVaultsServer() (pb.VaultsServer, error) {
	return &vaultsServer{&pb.VaultsDefaultServer{}}, nil
}

type applicationsServer struct {
	*pb.ApplicationsDefaultServer
}

// NewApplicationsServer returns an instance of the default Application server interface
func NewApplicationsServer() (pb.ApplicationsServer, error) {
	return &applicationsServer{&pb.ApplicationsDefaultServer{}}, nil
}

type appVersionsServer struct {
	*pb.AppVersionsDefaultServer
}

// NewAppVersionsServer returns an instance of the default AppVersion server interface
func NewAppVersionsServer() (pb.AppVersionsServer, error) {
	return &appVersionsServer{&pb.AppVersionsDefaultServer{}}, nil
}

type lifecyclesServer struct {
	*pb.LifecyclesDefaultServer
}

// NewLifecyclesServer returns an instance of the default Lifecycle server interface
func NewLifecyclesServer() (pb.LifecyclesServer, error) {
	return &lifecyclesServer{&pb.LifecyclesDefaultServer{}}, nil
}

type appConfigsServer struct {
	*pb.AppConfigsDefaultServer
}

// NewAppConfigsServer returns an instance of the default AppConfig server interface
func NewAppConfigsServer() (pb.AppConfigsServer, error) {
	return &appConfigsServer{&pb.AppConfigsDefaultServer{}}, nil
}

type chartVersionsServer struct {
	*pb.ChartVersionsDefaultServer
}

// NewChartVersionsServer returns an instance of the default ChartVersion server interface
func NewChartVersionsServer() (pb.ChartVersionsServer, error) {
	return &chartVersionsServer{&pb.ChartVersionsDefaultServer{}}, nil
}

type networksServer struct {
	*pb.NetworksDefaultServer
}
// NewNetworksServer returns an instance of the default Network server interface
func NewNetworksServer() (pb.NetworksServer, error) {
	return &networksServer{&pb.NetworksDefaultServer{}}, nil
}
