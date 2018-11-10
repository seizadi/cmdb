package svc

import (
	"github.com/seizadi/cmdb/pkg/pb"
)

type artifactsServer struct {
	*pb.ArtifactsDefaultServer
}
// NewArtifactsServer returns an instance of the default Artifact server interface
func NewArtifactsServer() (pb.ArtifactsServer, error) {
	return &artifactsServer{&pb.ArtifactsDefaultServer{}}, nil
}

type vaultsServer struct {
	*pb.VaultsDefaultServer
}
// NewVaultsServer returns an instance of the default Vault server interface
func NewVaultsServer() (pb.VaultsServer, error) {
	return &vaultsServer{&pb.VaultsDefaultServer{}}, nil
}

type versionTagsServer struct {
	*pb.VersionTagsDefaultServer
}
// NewVersionTagsServer returns an instance of the default VersionTag server interface
func NewVersionTagsServer() (pb.VersionTagsServer, error) {
	return &versionTagsServer{&pb.VersionTagsDefaultServer{}}, nil
}

type deploymentsServer struct {
	*pb.DeploymentsDefaultServer
}
// NewDeploymentsServer returns an instance of the default Deployment server interface
func NewDeploymentsServer() (pb.DeploymentsServer, error) {
	return &deploymentsServer{&pb.DeploymentsDefaultServer{}}, nil
}

type environmentsServer struct {
	*pb.EnvironmentsDefaultServer
}
// NewEnvironmentsServer returns an instance of the default Environment server interface
func NewEnvironmentsServer() (pb.EnvironmentsServer, error) {
	return &environmentsServer{&pb.EnvironmentsDefaultServer{}}, nil
}

type kubeClustersServer struct {
	*pb.KubeClustersDefaultServer
}
// NewKubeClustersServer returns an instance of the default KubeCluster server interface
func NewKubeClustersServer() (pb.KubeClustersServer, error) {
	return &kubeClustersServer{&pb.KubeClustersDefaultServer{}}, nil
}

type manifestsServer struct {
	*pb.ManifestsDefaultServer
}
// NewManifestsServer returns an instance of the default Manifest server interface
func NewManifestsServer() (pb.ManifestsServer, error) {
	return &manifestsServer{&pb.ManifestsDefaultServer{}}, nil
}

type applicationsServer struct {
	*pb.ApplicationsDefaultServer
}
// NewApplicationsServer returns an instance of the default Application server interface
func NewApplicationsServer() (pb.ApplicationsServer, error) {
	return &applicationsServer{&pb.ApplicationsDefaultServer{}}, nil
}

type awsRdsInstancesServer struct {
	*pb.AwsRdsInstancesDefaultServer
}
// NewAwsRdsInstancesServer returns an instance of the default AwsRdsInstance server interface
func NewAwsRdsInstancesServer() (pb.AwsRdsInstancesServer, error) {
	return &awsRdsInstancesServer{&pb.AwsRdsInstancesDefaultServer{}}, nil
}

type awsServicesServer struct {
	*pb.AwsServicesDefaultServer
}
// NewAwsServicesServer returns an instance of the default AwsService server interface
func NewAwsServicesServer() (pb.AwsServicesServer, error) {
	return &awsServicesServer{&pb.AwsServicesDefaultServer{}}, nil
}

type containersServer struct {
	*pb.ContainersDefaultServer
}
// NewContainersServer returns an instance of the default Container server interface
func NewContainersServer() (pb.ContainersServer, error) {
	return &containersServer{&pb.ContainersDefaultServer{}}, nil
}

type regionsServer struct {
	*pb.RegionsDefaultServer
}
// NewRegionsServer returns an instance of the default Region server interface
func NewRegionsServer() (pb.RegionsServer, error) {
	return &regionsServer{&pb.RegionsDefaultServer{}}, nil
}

type secretsServer struct {
	*pb.SecretsDefaultServer
}
// NewSecretsServer returns an instance of the default Secret server interface
func NewSecretsServer() (pb.SecretsServer, error) {
	return &secretsServer{&pb.SecretsDefaultServer{}}, nil
}

