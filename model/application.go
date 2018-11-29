package model

import "time"

type Application struct {
	ID           uint      `gorm:"primary_key" json:"id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	Name         string
	Description  string
	AppName      string
	Repo         string
	VersionTag   VersionTag `gorm:"foreignkey:VersionTagID"`
	VersionTagID uint
	Manifest     Manifest `gorm:"foreignkey:ManifestID"`
	ManifestID   uint
	Containers   []Container
	Deployment   Deployment
}


