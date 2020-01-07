package model

import "time"

type Value struct {
	ID           uint      `gorm:"primary_key" json:"id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	Name         string
	Description  string
	Keys          string     // jsonb
	AwsRdsInstance   AwsRdsInstance
	Secrets      []Secret
}


