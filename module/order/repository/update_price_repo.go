package orderRepo

import (
	"context"
	"go_service_food_organic/common"
	orderModel "go_service_food_organic/module/order/model"
	"reflect"
)

type UpdateOrderPriceStore interface {
	FindDataWithCondition(c context.Context, cond map[string]interface{}, oreKeys ...string) (*orderModel.Order, error)
	Update(c context.Context, id int, data *orderModel.Order) error
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
	data *orderModel.OrderUpdate) error {

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
	if order.State != "processing" {
		return common.ErrInternal(nil)
	}

	val := reflect.ValueOf(data).Elem()

	for i := 0; i < val.NumField(); i++ {
		field := val.Type().Field(i)
		value := val.Field(i).Interface()

		if value != "" {
			reflect.ValueOf(order).Elem().FieldByName(field.Name).Set(reflect.ValueOf(value))
		}
	}

	if err := repo.store.Update(c, order.Id, order); err != nil {
		return common.ErrCannotCRUDEntity(orderModel.EntityName, common.Update, err)
	}
	return nil
}
