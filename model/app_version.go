package model

import "time"

type AppVersion struct {
	ID           int       `gorm:"primary_key" json:"id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	Name         string
	Description  string
	ChartVersion ChartVersion
	Application  Application
}
