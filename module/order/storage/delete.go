package orderStorage

import (
	"context"
	"go_service_food_organic/common"
	orderModel "go_service_food_organic/module/order/model"
)

func (sql *sqlModel) DeleteOrder(c context.Context, idUser int) error {
	if err := sql.db.Table(orderModel.Order{}.TableName()).
		Where("id = (?)", idUser).
		Updates(map[string]interface{}{"status": 0}).
		Error; err != nil {
		return common.ErrDB(err)
	}
	return nil

}
