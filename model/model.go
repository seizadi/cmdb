package model

import "time"

// Override Gorm g.Model to add JSON annotations, should be in Gorm Model
type Model struct {
	ID        uint `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	//DeletedAt *time.Time `sql:"index;DEFAULT:0" json:"deleted_at"`
}
