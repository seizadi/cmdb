package model

type Manifest struct {
	Model
	Name         string
	Description  string
	Repo         string
	CommitId     string
	Values       string
	Services     string
	Ingress      string
	AwsService   AwsService `gorm:"foreignkey:AwsServiceID"`
	AwsServiceID uint
	Artifact     Artifact `gorm:"foreignkey:ArtifactID"`
	ArtifactID   uint
	Vault        Vault `gorm:"foreignkey:VaultID"`
	VaultID      uint
}

