package model

import "time"

type AppRegionConfig struct {
	ID              uint      `gorm:"primary_key" json:"id"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	Name            string
	Description     string
	ChartVersion    ChartVersion
	AppStageConfigs []AppStageConfig
	Value           Value
	Region          Region
}
