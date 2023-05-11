package orderDetailRepo

import (
	"context"
	"go_service_food_organic/common"
	orderDetailModel "go_service_food_organic/module/order_detail/model"
)

type ListOrderDetailStore interface {
	ListDataWithFilter(
		c context.Context,
		filter *orderDetailModel.Filter,
		paging *common.Paging,
		moreKeys ...string) ([]orderDetailModel.OrderDetail, error)
}

type listOrderDetailRepo struct {
	store ListOrderDetailStore
	req   common.Requester
}

func NewListOrderDetailRepo(store ListOrderDetailStore, req common.Requester) *listOrderDetailRepo {
	return &listOrderDetailRepo{store: store, req: req}
}

func (repo *listOrderDetailRepo) ListOrderDetailRepo(
	c context.Context,
	filter *orderDetailModel.Filter,
	paging *common.Paging) ([]orderDetailModel.OrderDetail, error) {

	list, err := repo.store.ListDataWithFilter(c, filter, paging)
	if err != nil || list == nil {
		return nil, common.ErrCannotCRUDEntity(orderDetailModel.EntityName, common.List, err)
	}
	return list, nil
}
