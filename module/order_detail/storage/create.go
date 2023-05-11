package orderDetailStorage

import (
	"context"
	"go_service_food_organic/common"
	orderDetailModel "go_service_food_organic/module/order_detail/model"
)

func (sql *sqlModel) Create(c context.Context, data *orderDetailModel.OrderDetailCreate) error {
	if err := sql.db.Table(orderDetailModel.OrderDetail{}.TableName()).Create(data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
