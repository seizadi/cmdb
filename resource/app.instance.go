package resource

import (
	"errors"
	"fmt"
	atlasresource "github.com/infobloxopen/atlas-app-toolkit/rpc/resource"
	"github.com/jinzhu/gorm"

	"github.com/seizadi/cmdb/pkg/pb"
)

var ErrUserNotFound = errors.New("app instance not found")

// GetAppInstanceById fetches an app instance from the database by id.
func GetAppInstanceById(id *atlasresource.Identifier, db *gorm.DB) (*pb.ApplicationInstanceORM, error) {
	appInstance := &pb.ApplicationInstanceORM{}
	err := db.Model(appInstance).First(appInstance, "id=?", id.ResourceId).Error
	if err != nil {
		return nil, fmt.Errorf("%s: %s", ErrUserNotFound, err.Error())
	}
	return appInstance, nil
}
