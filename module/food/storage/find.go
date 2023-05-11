package foodStorage

import (
	"context"
	"go_service_food_organic/common"
	foodModel "go_service_food_organic/module/food/model"
	"gorm.io/gorm"
)

func (sql *sqlModel) FindDataWithCondition(c context.Context, cond map[string]interface{}) (*foodModel.Food, error) {
	var food foodModel.Food
	if err := sql.db.Table(foodModel.Food{}.TableName()).Where(cond).First(&food).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.ErrRecordNotFound(foodModel.EntityName, err)
		}
		return nil, common.ErrDB(err)
	}
	return &food, nil
}
