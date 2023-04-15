package imageStorage

import (
	"context"
	"go_service_food_organic/common"
	"go_service_food_organic/module/upload/image/model"
	"gorm.io/gorm"
)

func (sql *sqlModel) FindDataWithCondition(c context.Context, cond map[string]interface{}) (*imageModel.Image, error) {
	var img imageModel.Image
	if err := sql.db.Table(imageModel.Image{}.TableName()).Where(cond).First(&img).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.ErrRecordNotFound(imageModel.EntityName, err)
		}
		return nil, common.ErrDB(err)
	}
	return &img, nil
}
