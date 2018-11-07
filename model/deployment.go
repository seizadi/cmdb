package model

type Deployment struct {
	Model
	Name         string
	Description  string
	Artifact     Artifact `gorm:"foreignkey:ArtifactID"`
	ArtifactID   uint
	Kubernetes   Kubernetes `gorm:"foreignkey:KubernetesID"`
	KubernetesID uint
}
