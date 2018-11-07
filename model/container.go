package model

type Container struct {
	Model
	Name            string
	ContainerName   string
	Description     string
	ImageRepo       string
	ImageTag        string
	ImagePullPolicy string
	Digest          string
}
