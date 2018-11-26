package model

import "time"

type AwsRdsInstance struct {
	ID        uint `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name             string
	Description      string
	DatabaseHost     string
	DatabaseName     string
	DatabaseUser     string
	DatabasePassword Secret
}




