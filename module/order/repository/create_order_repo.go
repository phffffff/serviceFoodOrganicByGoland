package orderRepo

import (
	"context"
	"go_service_food_organic/common"
	orderModel "go_service_food_organic/module/order/model"
)

type CreateOrderStore interface {
	Create(c context.Context, data *orderModel.OrderCreate) error
	BeginTransaction() error
	RollbackTransaction() error
	CommitTransaction() error
}
type createOrderRepo struct {
	store CreateOrderStore
	req   common.Requester
}

func NewCreateOrderRepo(store CreateOrderStore, req common.Requester) *createOrderRepo {
	return &createOrderRepo{store: store, req: req}
}

func (repo *createOrderRepo) CreateOrderRepo(c context.Context, data *orderModel.OrderCreate) error {
	if err := repo.store.Create(c, data); err != nil {
		return common.ErrCannotCRUDEntity(orderModel.EntityName, common.Create, err)
	}
	return nil
}
