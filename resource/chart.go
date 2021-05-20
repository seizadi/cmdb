package resource

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"

	"github.com/seizadi/cmdb/pkg/pb"
)

var ErrChartVersionNotFound = errors.New("chart version not found")

// GetChartVersionById fetches a chart version from the database by id.
func GetChartVersionById(id *string, db *gorm.DB) (*pb.ChartVersionORM, error) {
	charVersion := &pb.ChartVersionORM{}
	err := db.Model(charVersion).First(charVersion, "id=?", id).Error
	if err != nil {
		return nil, fmt.Errorf("%s: %s", ErrChartVersionNotFound, err.Error())
	}
	return charVersion, nil
}

