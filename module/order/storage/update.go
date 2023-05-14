package orderStorage

import (
	"context"
	"go_service_food_organic/common"
	orderModel "go_service_food_organic/module/order/model"
)

func (sql *sqlModel) Update(c context.Context, id int, data *orderModel.Order) error {
	if err := sql.db.Table(orderModel.Order{}.TableName()).Where("id = (?)", id).Updates(&data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
