package model

import "time"

type Deployment struct {
	ID            uint      `gorm:"primary_key" json:"id"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	Name          string
	Description   string
	Artifact      Artifact `gorm:"foreignkey:ArtifactID"`
	ArtifactID    uint
	KubeCluster   KubeCluster `gorm:"foreignkey:KubeClusterID"`
	KubeClusterID uint
}
