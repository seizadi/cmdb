package resource

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"

	"github.com/seizadi/cmdb/pkg/pb"
)

var ErrAppConfigNotFound = errors.New("appConfig instance not found")

// GetAppConfigByLifecycleId fetches an appConfig instance from the database by lifecycle id.
func GetAppConfigByLifecycleId(appId *string, lifecycleId *string, db *gorm.DB) (*pb.AppConfigORM, error) {
	appConfig := &pb.AppConfigORM{}
	err := db.Model(appConfig).First(appConfig, map[string]interface{}{"application_id": appId, "lifecycle_id": lifecycleId}).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, fmt.Errorf("%s: %s", ErrAppConfigNotFound, err.Error())
	}
	return appConfig, nil
}

// GetAppConfigByEnvId fetches an appConfig instance from the database by environment id.
func GetAppConfigByEnvId(appId *string, envId *string, db *gorm.DB) (*pb.AppConfigORM, error) {
	appConfig := &pb.AppConfigORM{}
	err := db.Model(appConfig).First(appConfig, map[string]interface{}{"application_id": appId, "environment_id": envId}).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, fmt.Errorf("%s: %s", ErrAppConfigNotFound, err.Error())
	}
	return appConfig, nil
}
