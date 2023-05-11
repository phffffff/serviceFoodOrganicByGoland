package orderDetailBusiness

import (
	"context"
	"go_service_food_organic/common"
	orderDetailModel "go_service_food_organic/module/order_detail/model"
)

type CreateOrderDetailRepo interface {
	CreateOrderDetailRepo(c context.Context, data *orderDetailModel.OrderDetailCreate) error
}

type createOrderDetailBiz struct {
	repo CreateOrderDetailRepo
}

func NewCreateOrderDetailBiz(repo CreateOrderDetailRepo) *createOrderDetailBiz {
	return &createOrderDetailBiz{repo: repo}
}

func (biz *createOrderDetailBiz) CreateOrderDetail(c context.Context, data *orderDetailModel.OrderDetailCreate) error {
	if err := biz.repo.CreateOrderDetailRepo(c, data); err != nil {
		return common.ErrInternal(err)
	}
	return nil
}
