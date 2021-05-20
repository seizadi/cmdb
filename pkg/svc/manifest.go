package svc

import (
	"context"
	"errors"
	"github.com/jinzhu/gorm"

	"github.com/seizadi/cmdb/helm"
	"github.com/seizadi/cmdb/pkg/pb"
	"github.com/seizadi/cmdb/resource"
)

type manifestServer struct {
	*pb.ManifestDefaultServer
}

var ErrAppInstanceDisabled = errors.New("app instance is disabled")

// NewManifestServer returns an instance of the default Manifest server interface
func NewManifestServer(db *gorm.DB) (pb.ManifestServer, error) {
	return &manifestServer{&pb.ManifestDefaultServer{db}}, nil
}

//ManifestCreate Creates a Manifest for a application instance
func (s *manifestServer) ManifestCreate(ctx context.Context, in *pb.ManifestCreateRequest) (*pb.ManifestCreateResponse, error) {
	db := s.DB
	response := pb.ManifestCreateResponse{}

	appInstance, err := resource.GetAppInstanceById(in.AppInstanceId, db)
	if err != nil {
		return &response, err
	}

	if appInstance.Enable == false {
		return &response, ErrAppInstanceDisabled
	}

	chartVersion, err := resource.GetChartVersionById(appInstance.ChartVersionId, db)
	if err != nil {
		return &response, err
	}

	h, err := helm.NewHelm()
	if err != nil {
		return &response, err
	}

	artifact, err := h.CreateManifest(chartVersion.Repo, chartVersion.Version)
	if err != nil {
		return &response, err
	}

	response.Artifact = artifact

	return &response, nil
}
