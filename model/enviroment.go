package model

import "time"

type Environment struct {
	ID        uint `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name string `json:"name"`
	Description string
	Code int
	Applications []Application `gorm:"many2many:environment_application";association_foreignkey:id;foreignkey:id`
}

