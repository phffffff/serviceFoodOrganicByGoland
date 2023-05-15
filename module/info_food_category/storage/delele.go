package infoFoodCategoryStorage

import (
	"context"
	"go_service_food_organic/common"
	infoFoodCategoryModel "go_service_food_organic/module/info_food_category/model"
)

func (sql *sqlModel) Delete(c context.Context, id int) error {
	if err := sql.db.
		Table(infoFoodCategoryModel.InfoFoodCategory{}.TableName()).
		Where("id = (?)", id).
		Updates(map[string]interface{}{"status": 0}).
		Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
