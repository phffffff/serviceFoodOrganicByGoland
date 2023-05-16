package brandStorage

import (
	"context"
	"go_service_food_organic/common"
	brandModel "go_service_food_organic/module/brand/model"
)

func (sql *sqlModel) Create(c context.Context, data *brandModel.BrandCreate) error {
	if err := sql.db.Table(brandModel.Brand{}.TableName()).Create(&data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
