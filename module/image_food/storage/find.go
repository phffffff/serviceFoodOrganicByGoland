package imageFoodStorage

import (
	"context"
	"go_service_food_organic/common"
	"go_service_food_organic/module/image_food/model"
	"gorm.io/gorm"
)

func (sql *sqlModel) FindDataWithCondition(c context.Context, cond map[string]interface{}) (*imageFoodModel.ImageFood, error) {
	var imgf imageFoodModel.ImageFood
	if err := sql.db.Table(imageFoodModel.ImageFood{}.TableName()).Where(cond).First(&imgf).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.ErrRecordNotFound(imageFoodModel.EntityName, err)
		}
		return nil, common.ErrDB(err)
	}
	return &imgf, nil
}
