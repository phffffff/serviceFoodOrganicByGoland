package orderDetailRepo

import (
	"context"
	"go_service_food_organic/common"
	foodModel "go_service_food_organic/module/food/model"
	orderDetailModel "go_service_food_organic/module/order_detail/model"
)

type CreateOrderDetailStore interface {
	Create(c context.Context, data *orderDetailModel.OrderDetailCreate) error
	BeginTransaction() error
	RollbackTransaction() error
	CommitTransaction() error
}

type FindFoodStore interface {
	FindDataWithCondition(c context.Context, cond map[string]interface{}) (*foodModel.Food, error)
}

type createOrderDetailRepo struct {
	store     CreateOrderDetailStore
	storeFood FindFoodStore
	req       common.Requester
}

func NewCreateOrderDetailRepo(store CreateOrderDetailStore, storeFood FindFoodStore, req common.Requester) *createOrderDetailRepo {
	return &createOrderDetailRepo{store: store, storeFood: storeFood, req: req}
}

func (repo *createOrderDetailRepo) CreateOrderDetailRepo(c context.Context, data *orderDetailModel.OrderDetailCreate) error {
	food, err := repo.storeFood.FindDataWithCondition(c, map[string]interface{}{"id": data.FoodId})
	if err != nil {
		return common.ErrRecordNotFound(foodModel.EntityName, err)
	}
	if food == nil {
		return common.ErrEntityNotExists(foodModel.EntityName, err)
	}
	if food.Status == 0 {
		return common.ErrEntityDeleted(foodModel.EntityName, nil)
	}

	if data.Quantity > food.Count {
		return orderDetailModel.ErrorQuantityInvalid(nil)
	}

	data.Price = food.Price

	if err := repo.store.Create(c, data); err != nil {
		return common.ErrCannotCRUDEntity(orderDetailModel.EntityName, common.Create, err)
	}
	return nil
}
