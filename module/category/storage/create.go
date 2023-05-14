package categoryStorage

import (
	"context"
	"go_service_food_organic/common"
	categoryModel "go_service_food_organic/module/category/model"
)

func (sql *sqlModel) Create(c context.Context, data *categoryModel.CategoryCreate) error {
	if err := sql.db.Table(categoryModel.Category{}.TableName()).Create(&data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
