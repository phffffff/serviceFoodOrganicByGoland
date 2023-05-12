package orderStorage

import (
	"context"
	"go_service_food_organic/common"
	orderModel "go_service_food_organic/module/order/model"
)

func (sql *sqlModel) Create(c context.Context, data *orderModel.OrderCreate) error {
	if err := sql.db.Table(orderModel.OrderCreate{}.TableName()).
		Create(&data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
