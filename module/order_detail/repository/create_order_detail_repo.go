package orderDetailRepo

import (
	"context"
	"go_service_food_organic/common"
	orderDetailModel "go_service_food_organic/module/order_detail/model"
)

type CreateOrderDetailStore interface {
	Create(c context.Context, data *orderDetailModel.OrderDetailCreate) error
}

type createOrderDetailRepo struct {
	store CreateOrderDetailStore
	req   common.Requester
}

func NewCreateOrderDetailRepo(store CreateOrderDetailStore, req common.Requester) *createOrderDetailRepo {
	return &createOrderDetailRepo{store: store, req: req}
}

func (repo *createOrderDetailRepo) CreateOrderDetailRepo(c context.Context, data *orderDetailModel.OrderDetailCreate) error {
	if err := repo.store.Create(c, data); err != nil {
		return common.ErrCannotCRUDEntity(orderDetailModel.EntityName, common.Create, err)
	}
	return nil
}
