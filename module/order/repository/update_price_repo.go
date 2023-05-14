package orderRepo

import (
	"context"
	"go_service_food_organic/common"
	orderModel "go_service_food_organic/module/order/model"
)

type UpdateOrderPriceStore interface {
	FindDataWithCondition(c context.Context, cond map[string]interface{}, moreKeys ...string) (*orderModel.Order, error)
	UpdatePrice(c context.Context, id int, price float32) error
}

type updateOrderPriceRepo struct {
	store UpdateOrderPriceStore
}

func NewUpdateOrderPriceRepo(store UpdateOrderPriceStore) *updateOrderPriceRepo {
	return &updateOrderPriceRepo{store: store}
}

func (repo *updateOrderPriceRepo) UpdateOrderPriceRepo(
	c context.Context,
	id int,
	price float32) error {

	order, err := repo.store.FindDataWithCondition(c, map[string]interface{}{"id": id})
	if err != nil {
		return common.ErrRecordNotFound(orderModel.EntityName, err)
	}
	if order == nil {
		return common.ErrEntityNotExists(orderModel.EntityName, err)
	}
	if order.Status == 0 {
		return common.ErrEntityDeleted(orderModel.EntityName, nil)
	}
	if order.State != orderModel.StateProcessing {
		return common.ErrInternal(nil)
	}

	if err := repo.store.UpdatePrice(c, order.Id, price); err != nil {
		return common.ErrCannotCRUDEntity(orderModel.EntityName, common.Update, err)
	}
	return nil
}
