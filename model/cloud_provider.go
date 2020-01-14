package model

import "time"

type CloudProvider struct {
	ID          int       `gorm:"primary_key" json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Name        string
	Provider    int
	Account     string
	Description string
	Regions     []Region
}
