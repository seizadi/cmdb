package model

type Vault struct {
	Model
	Name        string
	Description string
	Path        string
	Secrets     []Secret `gorm:"many2many:vault_secret;association_foreignkey:id;foreignkey:id"`
}
