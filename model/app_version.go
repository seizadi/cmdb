package model

import "time"

type AppVersion struct {
	ID          uint      `gorm:"primary_key" json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Name        string
	Application Application
	Description string
	Repo        string
	Version     string
}
