package foodStorage

import (
	"context"
	"go_service_food_organic/common"
	foodModel "go_service_food_organic/module/food/model"
	"gorm.io/gorm"
)

func (sql *sqlModel) Update(c context.Context, data *foodModel.Food, id int) error {
	if err := sql.db.Table(foodModel.Food{}.TableName()).
		Where("id = (?)", id).
		Updates(data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}

func (sql *sqlModel) UpdateCount(c context.Context, count int, id int) error {
	if err := sql.db.Table(foodModel.Food{}.TableName()).Where("id = (?)", id).
		Update("count", gorm.Expr("count - ?", count)).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
