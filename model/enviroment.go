package model

type Environment struct {
	Model
	Name string `json:"name"`
	Description string
	Code string
	Applications []Application `gorm:"many2many:environment_application";association_foreignkey:id;foreignkey:id`
}

