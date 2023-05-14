package foodStorage

import (
	"context"
	"fmt"
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

func (sql *sqlModel) UpdateCount(c context.Context, count int, id int, typeOf string) error {
	if err := sql.db.Table(foodModel.Food{}.TableName()).Where("id = (?)", id).
		Update("count", gorm.Expr(fmt.Sprintf("count %s ?", typeOf), count)).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
