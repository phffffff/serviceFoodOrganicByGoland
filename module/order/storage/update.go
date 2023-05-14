package orderStorage

import (
	"context"
	"go_service_food_organic/common"
	orderModel "go_service_food_organic/module/order/model"
)

func (sql *sqlModel) UpdateState(c context.Context, id int, state string) error {
	if state == orderModel.StateCancel {
		if err := sql.db.
			Table(orderModel.Order{}.TableName()).
			Where("id = (?)", id).
			Updates(map[string]interface{}{"status": 0, "state": orderModel.StateCancel}).Error; err != nil {
			return common.ErrDB(err)
		}
		return nil
	}
	if err := sql.db.Table(orderModel.Order{}.TableName()).
		Where("id = (?)", id).
		Updates(map[string]interface{}{"status": 1, "state": state}).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}

func (sql *sqlModel) UpdatePrice(c context.Context, id int, price float32) error {
	if err := sql.db.Table(orderModel.Order{}.TableName()).Where("id = (?)", id).Update("total_price", price).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}

func (sql *sqlModel) Update(c context.Context, id int, data *orderModel.OrderUpdate) error {
	if err := sql.db.Table(orderModel.OrderUpdate{}.TableName()).Where("id = (?)", id).Updates(&data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
