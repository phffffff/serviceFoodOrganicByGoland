package orderBusiness

import (
	"context"
	"go_service_food_organic/common"
	orderModel "go_service_food_organic/module/order/model"
)

type ListOrderRepo interface {
	ListOrderRepo(
		c context.Context,
		filter *orderModel.Filter,
		paging *common.Paging) ([]orderModel.Order, error)
}

type listOrderBiz struct {
	repo ListOrderRepo
}

func NewListOrderBiz(repo ListOrderRepo) *listOrderBiz {
	return &listOrderBiz{repo: repo}
}

func (biz *listOrderBiz) ListOrder(
	c context.Context,
	filter *orderModel.Filter,
	paging *common.Paging) ([]orderModel.Order, error) {

	list, err := biz.repo.ListOrderRepo(c, filter, paging)
	if err != nil || list == nil {
		return nil, common.ErrCannotCRUDEntity(orderModel.EntityName, common.List, err)
	}
	return list, nil
}
