package model

import "time"

type ApplicationInstance struct {
	ID           uint      `gorm:"primary_key" json:"id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	Name         string
	Description  string
	AppName      string
	Repo         string
	ChartVersion   ChartVersion `gorm:"foreignkey:ChartVersionID"`
	VersionTagID uint
	Values     []Value
	Deployment   Deployment
}


