package infoFoodCategoryStorage

import (
	"context"
	"go_service_food_organic/common"
	infoFoodCategoryModel "go_service_food_organic/module/info_food_category/model"
	"gorm.io/gorm"
)

func (sql *sqlModel) FindDataWithCondition(c context.Context, cond map[string]interface{}) (*infoFoodCategoryModel.InfoFoodCategory, error) {
	var ifc infoFoodCategoryModel.InfoFoodCategory
	if err := sql.db.Table(infoFoodCategoryModel.InfoFoodCategory{}.TableName()).Where(cond).First(&ifc).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.ErrRecordNotFound(infoFoodCategoryModel.EntityName, err)
		}
		return nil, common.ErrDB(err)
	}
	return &ifc, nil
}
