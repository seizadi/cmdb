package svc

import (
	"context"
	"github.com/jinzhu/gorm"
	"github.com/seizadi/cmdb/pkg/pb"
)

type manifestServer struct {
	*pb.ManifestDefaultServer
}

// NewManifestServer returns an instance of the default Manifest server interface
func NewManifestServer(db *gorm.DB) (pb.ManifestServer, error) {
	return &manifestServer{&pb.ManifestDefaultServer{db}}, nil
}

//ManifestCreate Creates a Manifest for a application instance
func (s *manifestServer) ManifestCreate(ctx context.Context, in *pb.ManifestCreateRequest) (*pb.ManifestCreateResponse, error) {
	response := pb.ManifestCreateResponse{}
	return &response, nil
}
