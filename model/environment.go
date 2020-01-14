package model

import "time"

type Environment struct {
	ID                   int       `gorm:"primary_key" json:"id"`
	CreatedAt            time.Time `json:"created_at"`
	UpdatedAt            time.Time `json:"updated_at"`
	Name                 string    `json:"name"`
	Description          string
	AppConfigs           []AppConfig
	AppVersions          []AppVersion
	ApplicationInstances []ApplicationInstance
}
