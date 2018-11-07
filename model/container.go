package model

type Container struct {
	Model
	Name            string
	Description     string
	ContainerName   string
	ImageRepo       string
	ImageTag        string
	ImagePullPolicy string
	Digest          string
}
