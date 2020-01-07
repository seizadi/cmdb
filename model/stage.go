package model

import "time"

type Stage struct {
	ID           uint      `gorm:"primary_key" json:"id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	Name         string
	Description  string
	RegionID     uint
	Value      Value
	Applications []Application
	Environments []Environment
}
