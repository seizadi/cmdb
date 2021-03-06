package svc

import (
	"context"
	"errors"
	"strings"
	
	"github.com/jinzhu/gorm"
	"go.uber.org/config"
	"gopkg.in/yaml.v2"
	
	"github.com/seizadi/cmdb/helm"
	"github.com/seizadi/cmdb/pkg/pb"
	"github.com/seizadi/cmdb/resource"
	"github.com/seizadi/cmdb/utils"
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

	config, err := s.getManifestConfig(appInstance)
	if err != nil {
		return &response, err
	}

	configPath, err := utils.WriteTempFile(config, "config")
	if err != nil {
		return &response, err
	}


	artifact, err := h.CreateManifest(chartVersion.Repo, chartVersion.Version, configPath)
	if err != nil {
		return &response, err
	}

	response.Artifact = artifact

	return &response, nil
}

func (s *manifestServer) getManifestConfig(appInstance *pb.ApplicationInstanceORM) (string, error) {
	db := s.DB

	// Get Environment resource
	environment, err := resource.GetEnvrionmentById(appInstance.EnvironmentId, db)
	if err != nil {
		return "", err
	}

	// Get Lifecycle Resource until we reach the root
	var lifecycles []*pb.LifecycleORM

	lifecycle, err := resource.GetLifecycleById(environment.LifecycleId, db)
	if err != nil {
		return "", err
	}

	lifecycles  = append([]*pb.LifecycleORM{lifecycle}, lifecycles...)
	LifecycleId := lifecycle.LifecycleId

	// TODO - I did not put a count here to terminate and prevent infinite loop
	for {
		if LifecycleId == nil { // Reached the root node
			break
		}

		lifecycle, err := resource.GetLifecycleById(lifecycle.LifecycleId, db)
		if err != nil {
			return "", err
		}
		LifecycleId = lifecycle.LifecycleId

		lifecycles  = append([]*pb.LifecycleORM{lifecycle}, lifecycles...)
	}

	// Now fetch the lifecycle, env and app configuration values
	// Values are applied in this order: app -> env (app -> value) -> lifecycle (app -> value) -> lifecycle (app -> value)
	// The app value has highest precedent and over-rides lower values, environment and
	// lifecycle have app specific config and default value for all apps.

	var v map[interface{}]interface{}

	// One prolblem with this pattern is that errors are not detected until we merge using config.NewYAML
	var sources []config.YAMLOption
	var source config.YAMLOption

	// For each lifecycle in the list get the values and app configuration
	for _,l := range lifecycles {
		source = config.Source(strings.NewReader(l.ConfigYaml))
		//sources = append([]config.YAMLOption{source}, sources...)
		sources = append(sources, source)

		// Find the AppConfig
		appConfig, err := resource.GetAppConfigByLifecycleId(appInstance.ApplicationId, &l.Id, db)
		if err != nil {
			return "", err
		}

		if appConfig != nil {
			source = config.Source(strings.NewReader(appConfig.ConfigYaml))
			//sources = append([]config.YAMLOption{source}, sources...)
			sources = append(sources, source)		}
	}

	// Now mix in the environment config
	source = config.Source(strings.NewReader(environment.ConfigYaml))
	//sources = append([]config.YAMLOption{source}, sources...)
	sources = append(sources, source)

	// Find the environment AppConfig
	appConfig, err := resource.GetAppConfigByEnvId(appInstance.ApplicationId, &environment.Id, db)
	if err != nil {
		return "", err
	}

	if appConfig != nil {
		source = config.Source(strings.NewReader(appConfig.ConfigYaml))
		//sources = append([]config.YAMLOption{source}, sources...)
		sources = append(sources, source)
	}

	// Mix in the application config
	source = config.Source(strings.NewReader(appInstance.ConfigYaml))
	//sources = append([]config.YAMLOption{source}, sources...)
	sources = append(sources, source)
	sources = append(sources, config.Permissive())

	provider, err := config.NewYAML(sources...)
	if err != nil {
		return "", err
	}

	err = provider.Get(config.Root).Populate(&v)

	// Create the Yaml File
	c, err := yaml.Marshal(&v)
	if err != nil {
		return "", err
	}

	// Originally I was using helm to resolve the values, it was taking about 370ms
	// I wrote it to sue go template engine directly which reduced the time around 70ms
	values := helm.Values{Values: utils.MergeInterfaceToStringMaps(v)}
	r := helm.Renderable { Tpl: string(c), Vals: values}
	config, err := helm.RenderWithReferences(r)
	if err != nil {
		return "", err
	}

	return config, nil
}

//ManifestConfigCreate Creates a Config for a application instance
func (s *manifestServer) ManifestConfigCreate(ctx context.Context, in *pb.ManifestConfigCreateRequest) (*pb.ManifestConfigCreateResponse, error) {
	db := s.DB
	response := pb.ManifestConfigCreateResponse{}

	appInstance, err := resource.GetAppInstanceById(in.AppInstanceId, db)
	if err != nil {
		return &response, err
	}

	config, err := s.getManifestConfig(appInstance)
	if err != nil {
		return &response, err
	}

	response.Config = config
	return &response, nil
}
