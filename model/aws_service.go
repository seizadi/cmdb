package model

type AwsService struct {
	Model
	Name string
	Description string
	AwsRds []AwsRds `gorm:"many2many:aws_to_rds;association_foreignkey:id;foreignkey:id"`
}
