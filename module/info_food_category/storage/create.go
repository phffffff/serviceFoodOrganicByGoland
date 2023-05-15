package infoFoodCategoryStorage

import (
	"context"
	"go_service_food_organic/common"
	infoFoodCategoryModel "go_service_food_organic/module/info_food_category/model"
)

func (sql *sqlModel) Create(c context.Context, data *infoFoodCategoryModel.InfoFoodCategoryCreate) error {
	if err := sql.db.Table(infoFoodCategoryModel.InfoFoodCategoryCreate{}.TableName()).Create(&data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
