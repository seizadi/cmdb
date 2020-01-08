package model

import "time"

type ApplicationCatalog struct {
	ID           uint      `gorm:"primary_key" json:"id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	Name         string
	Description  string
	ChartVersion ChartVersion
	Applications []Application
	Value        Value
	Region       Region
}
