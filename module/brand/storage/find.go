package brandStorage

import (
	"context"
	"go_service_food_organic/common"
	brandModel "go_service_food_organic/module/brand/model"
	"gorm.io/gorm"
)

func (sql *sqlModel) FindDataWithCondition(c context.Context, cond map[string]interface{}) (*brandModel.Brand, error) {
	var brand brandModel.Brand
	if err := sql.db.Table(brandModel.Brand{}.TableName()).Where(cond).First(&brand).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.ErrRecordNotFound(brandModel.EntityName, err)
		}
		return nil, common.ErrDB(err)
	}
	return &brand, nil
}
