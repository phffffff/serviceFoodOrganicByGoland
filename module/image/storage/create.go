package imageStorage

import (
	"context"
	"go_service_food_organic/common"
	"go_service_food_organic/module/image/model"
)

func (sql *sqlModel) Create(c context.Context, data *imageModel.Image) error {
	if err := sql.db.Table(imageModel.Image{}.TableName()).Create(&data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
