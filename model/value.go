package model

import "time"

type Value struct {
	ID           uint      `gorm:"primary_key" json:"id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	Name         string
	Description  string
	Key          string     // jsonb
	AwsService   AwsService `gorm:"foreignkey:AwsServiceID"`
	AwsServiceID uint
	Secrets      []Secret
}

