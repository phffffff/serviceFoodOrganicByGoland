package categoryStorage

import (
	"context"
	"go_service_food_organic/common"
	categoryModel "go_service_food_organic/module/category/model"
)

func (sql *sqlModel) Delete(c context.Context, id int) error {
	if err := sql.db.Table(categoryModel.Category{}.TableName()).Where("id = (?)", id).
		Updates(map[string]interface{}{"status": 0}).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
