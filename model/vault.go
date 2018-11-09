package model

import "time"

type Vault struct {
	ID        uint `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name        string
	Description string
	Path        string
	Secrets     []Secret `gorm:"many2many:vault_secret;association_foreignkey:id;foreignkey:id"`
}
