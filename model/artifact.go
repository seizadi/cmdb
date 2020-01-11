package model

import "time"

type Artifact struct {
	ID          uint      `gorm:"primary_key" json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Name        string
	Description string
	Repo        string
	Commit      string
	AppVersion  AppVersion
}
