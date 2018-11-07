package model

type Region struct {
	Model
	Name string
	Description string
	Account string
	Environments []Environment `gorm:"many2many:region_environment";association_foreignkey:id;foreignkey:id`
}

