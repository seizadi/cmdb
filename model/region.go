package model

import "time"

type Region struct {
	ID        uint `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name string
	Description string
	Environments []Environment `gorm:"many2many:region_environment";association_foreignkey:id;foreignkey:id`
}

