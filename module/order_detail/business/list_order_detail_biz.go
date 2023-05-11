package orderDetailBusiness

import (
	"context"
	"go_service_food_organic/common"
	orderDetailModel "go_service_food_organic/module/order_detail/model"
)

type ListOrderDetailRepo interface {
	ListOrderDetailRepo(
		c context.Context,
		filter *orderDetailModel.Filter,
		paging *common.Paging,
	) ([]orderDetailModel.OrderDetail, error)
}

type listOrderDetailBiz struct {
	repo ListOrderDetailRepo
}

func NewListOrderDetailBiz(repo ListOrderDetailRepo) *listOrderDetailBiz {
	return &listOrderDetailBiz{repo: repo}
}

func (biz *listOrderDetailBiz) ListOrderDetail(
	c context.Context,
	filter *orderDetailModel.Filter,
	paging *common.Paging) ([]orderDetailModel.OrderDetail, error) {

	list, err := biz.repo.ListOrderDetailRepo(c, filter, paging)
	if err != nil || list == nil {
		return nil, common.ErrCannotCRUDEntity(orderDetailModel.EntityName, common.List, err)
	}
	return list, nil
}
