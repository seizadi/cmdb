package model

import "time"

type Lifecycle struct {
	Lifecycles []Lifecycle
	ID          uint      `gorm:"primary_key" json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Name        string
	Description string
	ConfigYaml  string
	AppConfigs  []AppConfig
	AppVersions []AppVersion
	Environment []Environment
}
