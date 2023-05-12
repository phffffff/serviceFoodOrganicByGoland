package orderRepo

import (
	"context"
	"go_service_food_organic/common"
	orderModel "go_service_food_organic/module/order/model"
)

type ListOrderStore interface {
	ListDataWithFilter(
		c context.Context,
		filter *orderModel.Filter,
		paging *common.Paging,
		moreKeys ...string) ([]orderModel.Order, error)
}

type listOrderRepo struct {
	store ListOrderStore
	req   common.Requester
}

func NewListOrderRepo(store ListOrderStore, req common.Requester) *listOrderRepo {
	return &listOrderRepo{store: store, req: req}
}

func (repo *listOrderRepo) ListOrderRepo(
	c context.Context,
	filter *orderModel.Filter,
	paging *common.Paging) ([]orderModel.Order, error) {

	list, err := repo.store.ListDataWithFilter(c, filter, paging, "OrderDetails.Foods.FoodImages", "Users")
	if err != nil || list == nil {
		return nil, common.ErrCannotCRUDEntity(orderModel.EntityName, common.List, err)
	}
	return list, nil
}
