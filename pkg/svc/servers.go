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

type stagesServer struct {
	*pb.StagesDefaultServer
}
// NewStagesServer returns an instance of the default Stage server interface
func NewStagesServer() (pb.StagesServer, error) {
	return &stagesServer{&pb.StagesDefaultServer{}}, nil
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

type applicationsServer struct {
	*pb.ApplicationsDefaultServer
}
// NewApplicationsServer returns an instance of the default Application server interface
func NewApplicationsServer() (pb.ApplicationsServer, error) {
	return &applicationsServer{&pb.ApplicationsDefaultServer{}}, nil
}

type kubeClustersServer struct {
	*pb.KubeClustersDefaultServer
}
// NewKubeClustersServer returns an instance of the default KubeCluster server interface
func NewKubeClustersServer() (pb.KubeClustersServer, error) {
	return &kubeClustersServer{&pb.KubeClustersDefaultServer{}}, nil
}

type awsRdsInstancesServer struct {
	*pb.AwsRdsInstancesDefaultServer
}
// NewAwsRdsInstancesServer returns an instance of the default AwsRdsInstance server interface
func NewAwsRdsInstancesServer() (pb.AwsRdsInstancesServer, error) {
	return &awsRdsInstancesServer{&pb.AwsRdsInstancesDefaultServer{}}, nil
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

type valuesServer struct {
	*pb.ValuesDefaultServer
}
// NewValuesServer returns an instance of the default Value server interface
func NewValuesServer() (pb.ValuesServer, error) {
	return &valuesServer{&pb.ValuesDefaultServer{}}, nil
}

type awsServicesServer struct {
	*pb.AwsServicesDefaultServer
}
// NewAwsServicesServer returns an instance of the default AwsService server interface
func NewAwsServicesServer() (pb.AwsServicesServer, error) {
	return &awsServicesServer{&pb.AwsServicesDefaultServer{}}, nil
}

type chartVersionsServer struct {
	*pb.ChartVersionsDefaultServer
}
// NewChartVersionsServer returns an instance of the default ChartVersion server interface
func NewChartVersionsServer() (pb.ChartVersionsServer, error) {
	return &chartVersionsServer{&pb.ChartVersionsDefaultServer{}}, nil
}
