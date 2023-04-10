package imageStorage

import (
	"context"
	"go_service_food_organic/common"
	"gorm.io/gorm"
)

func (sql *sqlModel) FindDataWithCondition(c context.Context, cond map[string]interface{}) (*common.Image, error) {
	var img common.Image
	if err := sql.db.Table(common.Image{}.GetTableName()).Where(cond).First(&img).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.ErrRecordNotFound(common.EntityName, err)
		}
		return nil, common.ErrDB(err)
	}
	return &img, nil
}
