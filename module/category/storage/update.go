package categoryStorage

import (
	"context"
	"go_service_food_organic/common"
	categoryModel "go_service_food_organic/module/category/model"
)

func (sql *sqlModel) Update(c context.Context, data *categoryModel.Category, id int) error {
	if err := sql.db.Table(categoryModel.Category{}.TableName()).
		Where("id = (?)", id).
		Updates(data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
