package model

import "time"

type AwsService struct {
	ID        uint `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name string
	Description string
	AwsRdsInstance []AwsRdsInstance `gorm:"many2many:aws_to_rds;association_foreignkey:id;foreignkey:id"`
}
