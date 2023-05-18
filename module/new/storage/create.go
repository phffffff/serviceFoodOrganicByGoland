package newStorage

import (
	"context"
	"go_service_food_organic/common"
	newModel "go_service_food_organic/module/new/model"
)

func (sql *sqlModel) Create(c context.Context, data *newModel.NewCrt) error {
	if err := sql.db.Table(newModel.NewCrt{}.TableName()).Create(&data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
