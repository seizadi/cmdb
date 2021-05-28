package resource

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"

	"github.com/seizadi/cmdb/pkg/pb"
)

var ErrEnvironmentNotFound = errors.New("environment instance not found")

// GetEnvrionmentById fetches an environment instance from the database by id.
func GetEnvrionmentById(id *string, db *gorm.DB) (*pb.EnvironmentORM, error) {
	environment := &pb.EnvironmentORM{}
	err := db.Model(environment).First(environment, "id=?", id).Error
	if err != nil {
		return nil, fmt.Errorf("%s: %s", ErrEnvironmentNotFound, err.Error())
	}
	return environment, nil
}
