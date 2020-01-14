package model

import "time"

type ApplicationInstance struct {
	ID          int       `gorm:"primary_key" json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Name        string
	Description string
	AppVersion  AppVersion
	Secrets     []Secret
	Enable      bool
	Deployment  Deployment
}
