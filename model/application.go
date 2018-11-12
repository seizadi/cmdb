package model

import "time"

type Application struct {
	ID        uint `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name         string
	Description  string
	VersionTag   VersionTag `gorm:"foreignkey:VersionTagID"`
	VersionTagID uint
	Manifest     Manifest    `gorm:"foreignkey:ManifestID"`
	ManifestID    uint
	Containers []Container `gorm:"many2many:application_container";association_foreignkey:id;foreignkey:id`
}


