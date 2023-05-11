package foodStorage

import (
	"context"
	"go_service_food_organic/common"
	foodModel "go_service_food_organic/module/food/model"
)

func (sql *sqlModel) Delete(c context.Context, id int) error {
	if err := sql.db.Table(foodModel.Food{}.TableName()).Where("id = (?)", id).
		Updates(map[string]interface{}{"status": 0}).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
