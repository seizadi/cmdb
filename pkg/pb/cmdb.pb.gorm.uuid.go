package pb

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

func (pOrm CloudProviderORM) BeforeSave(scope *gorm.Scope) error {
	if pOrm.Id == "" {
		return scope.SetColumn("id", uuid.NewV4().String())
	}
	return nil
}

func (pOrm RegionORM) BeforeSave(scope *gorm.Scope) error {
	if pOrm.Id == "" {
		return scope.SetColumn("id", uuid.NewV4().String())
	}
	return nil
}

func (pOrm NetworkORM) BeforeSave(scope *gorm.Scope) error {
	if pOrm.Id == "" {
		return scope.SetColumn("id", uuid.NewV4().String())
	}
	return nil
}

func (pOrm LifecycleORM) BeforeSave(scope *gorm.Scope) error {
	if pOrm.Id == "" {
		return scope.SetColumn("id", uuid.NewV4().String())
	}
	return nil
}

func (pOrm EnvironmentORM) BeforeSave(scope *gorm.Scope) error {
	if pOrm.Id == "" {
		return scope.SetColumn("id", uuid.NewV4().String())
	}
	return nil
}

func (pOrm ApplicationORM) BeforeSave(scope *gorm.Scope) error {
	if pOrm.Id == "" {
		return scope.SetColumn("id", uuid.NewV4().String())
	}
	return nil
}

func (pOrm ChartVersionORM) BeforeSave(scope *gorm.Scope) error {
	if pOrm.Id == "" {
		return scope.SetColumn("id", uuid.NewV4().String())
	}
	return nil
}

func (pOrm AppVersionORM) BeforeSave(scope *gorm.Scope) error {
	if pOrm.Id == "" {
		return scope.SetColumn("id", uuid.NewV4().String())
	}
	return nil
}

func (pOrm AppConfigORM) BeforeSave(scope *gorm.Scope) error {
	if pOrm.Id == "" {
		return scope.SetColumn("id", uuid.NewV4().String())
	}
	return nil
}

func (pOrm ApplicationInstanceORM) BeforeSave(scope *gorm.Scope) error {
	if pOrm.Id == "" {
		return scope.SetColumn("id", uuid.NewV4().String())
	}
	return nil
}

func (pOrm VaultORM) BeforeSave(scope *gorm.Scope) error {
	if pOrm.Id == "" {
		return scope.SetColumn("id", uuid.NewV4().String())
	}
	return nil
}

func (pOrm SecretORM) BeforeSave(scope *gorm.Scope) error {
	if pOrm.Id == "" {
		return scope.SetColumn("id", uuid.NewV4().String())
	}
	return nil
}

func (pOrm ArtifactORM) BeforeSave(scope *gorm.Scope) error {
	if pOrm.Id == "" {
		return scope.SetColumn("id", uuid.NewV4().String())
	}
	return nil
}

func (pOrm KubeClusterORM) BeforeSave(scope *gorm.Scope) error {
	if pOrm.Id == "" {
		return scope.SetColumn("id", uuid.NewV4().String())
	}
	return nil
}

func (pOrm DeploymentORM) BeforeSave(scope *gorm.Scope) error {
	if pOrm.Id == "" {
		return scope.SetColumn("id", uuid.NewV4().String())
	}
	return nil
}
