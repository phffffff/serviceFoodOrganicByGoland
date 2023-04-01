package foodStorage

import (
	"context"
	"go_service_food_organic/common"
	foodModel "go_service_food_organic/module/food/model"
)

func (sql *sqlModel) Create(c context.Context, data *foodModel.FoodCreate) error {
	if err := sql.db.Table(foodModel.FoodCreate{}.GetTableName()).Create(&data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
