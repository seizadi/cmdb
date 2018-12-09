package model

import "time"

type Manifest struct {
	ID           uint      `gorm:"primary_key" json:"id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	Name         string
	Description  string
	Repo         string
	Commit       string
	Values       string
	Services     string
	Ingress      string
	Resources    string
	NodeSelector string
	Tolerations  string
	Affinity     string
	AwsService   AwsService `gorm:"foreignkey:AwsServiceID"`
	AwsServiceID uint
	Artifact     Artifact `gorm:"foreignkey:ArtifactID"`
	ArtifactID   uint
	Vault        Vault `gorm:"foreignkey:VaultID"`
	VaultID      uint
}


