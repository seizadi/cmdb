package resource

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"

	"github.com/seizadi/cmdb/pkg/pb"
)

var ErrLifecycleNotFound = errors.New("lifecycle instance not found")

// GetLifecycleById fetches a lifecycle instance from the database by id.
func GetLifecycleById(id *string, db *gorm.DB) (*pb.LifecycleORM, error) {
	lifecycle := &pb.LifecycleORM{}
	err := db.Model(lifecycle).First(lifecycle, "id=?", id).Error
	if err != nil {
		return nil, fmt.Errorf("%s: %s", ErrLifecycleNotFound, err.Error())
	}
	return lifecycle, nil
}
