package model

import "time"

type Container struct {
	ID        uint `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name            string
	Description     string
	ContainerName   string
	ImageRepo       string
	ImageTag        string
	ImagePullPolicy string
	Digest          string
}
