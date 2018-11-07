package model

type AwsRds struct {
	Model
	Name             string
	Description      string
	DatabaseHost     string
	DatabaseName     string
	DatabaseUser     string
	DatabasePassword Secret `gorm:"foreignkey:AwsRdsSecretID"`
	AwsRdsSecretID   uint
}




