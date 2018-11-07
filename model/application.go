package model

type Application struct {
	Model
	Name         string
	Description  string
	Code         string
	VersionTag   VersionTag `gorm:"foreignkey:VersionTagID"`
	VersionTagID uint
	Manifest     Manifest    `gorm:"foreignkey:ManifestID"`
	ManifestID    uint
	Containers []Container `gorm:"many2many:application_container";association_foreignkey:id;foreignkey:id`
}


