package categoryStorage

import (
	"context"
	"go_service_food_organic/common"
	categoryModel "go_service_food_organic/module/category/model"
	"gorm.io/gorm"
)

func (sql *sqlModel) FindDataWithCondition(c context.Context, cond map[string]interface{}) (*categoryModel.Category, error) {
	var category categoryModel.Category
	if err := sql.db.Table(categoryModel.Category{}.TableName()).Where(cond).First(&category).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.ErrRecordNotFound(categoryModel.EntityName, err)
		}
		return nil, common.ErrDB(err)
	}
	return &category, nil
}
