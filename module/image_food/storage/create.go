package imageFoodStorage

import (
	"context"
	"go_service_food_organic/common"
	"go_service_food_organic/module/image_food/model"
)

func (sql *sqlModel) Create(c context.Context, data *imageFoodModel.ImageFoodCreate) error {
	if err := sql.db.Table(imageFoodModel.ImageFood{}.TableName()).Create(&data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
