package model

import "time"

type KubeCluster struct {
	ID          int       `gorm:"primary_key" json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Name        string
	Description string
	Network     Network
}
