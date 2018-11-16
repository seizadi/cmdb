package model

import "time"

type CloudProvider struct {
	ID          uint      `gorm:"primary_key" json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Name        string
	Provider    string
	Account     string
	Description string
	Regions     []Region `gorm:"foreignkey:CloudProviderID;association_foreignkey:id"`
}
