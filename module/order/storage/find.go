package orderStorage

import (
	"context"
	"go_service_food_organic/common"
	orderModel "go_service_food_organic/module/order/model"
)

func (sql *sqlModel) FindDataWithCondition(
	c context.Context,
	cond map[string]interface{},
	moreKeys ...string) (*orderModel.Order, error) {

	db := sql.db.Table(orderModel.Order{}.TableName())
	if err := db.Error; err != nil {
		return nil, common.ErrDB(err)
	}

	var data orderModel.Order
	if err := db.Where(cond).First(&data).Error; err != nil {
		return nil, common.ErrDB(err)
	}
	return &data, nil
}
