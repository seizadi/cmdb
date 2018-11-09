package svc

import (
	"github.com/seizadi/cmdb/pkg/pb"
)

type regionsServer struct {
	*pb.RegionsDefaultServer
}
// NewRegionsServer returns an instance of the default regions server interface
func NewRegionsServer() (pb.RegionsServer, error) {
	return &regionsServer{&pb.RegionsDefaultServer{}}, nil
}

type containersServer struct {
	*pb.ContainersDefaultServer
}

// NewContainersServer returns an instance of the default regions server interface
func NewContainersServer() (pb.ContainersServer, error) {
	return &containersServer{&pb.ContainersDefaultServer{}}, nil
}

type versionTagsServer struct {
	*pb.VersionTagsDefaultServer
}

// NewVersionTagsServer returns an instance of the default regions server interface
func NewVersionTagsServer() (pb.VersionTagsServer, error) {
	return &versionTagsServer{&pb.VersionTagsDefaultServer{}}, nil
}

type secretsServer struct {
	*pb.SecretsDefaultServer
}

// NewSecretsServer returns an instance of the default regions server interface
func NewSecretsServer() (pb.SecretsServer, error) {
	return &secretsServer{&pb.SecretsDefaultServer{}}, nil
}

type vaultsServer struct {
	*pb.VaultsDefaultServer
}

// NewVaultsServer returns an instance of the default regions server interface
func NewVaultsServer() (pb.VaultsServer, error) {
	return &vaultsServer{&pb.VaultsDefaultServer{}}, nil
}
