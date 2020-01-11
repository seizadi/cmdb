package model

import "time"

type Secret struct {
	ID          uint      `gorm:"primary_key" json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Name        string
	Description string
	Vault       Vault `gorm:"foreignkey:VaultID"`
	VaultID     uint
	Type        string
	Key         string // jsonb
}
