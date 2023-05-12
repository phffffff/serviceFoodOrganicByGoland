package orderBusiness

import (
	"context"
	"go_service_food_organic/common"
	orderModel "go_service_food_organic/module/order/model"
)

type CreateOrderRepo interface {
	CreateOrderRepo(c context.Context, data *orderModel.OrderCreate) error
}

type createOrderBiz struct {
	repo CreateOrderRepo
}

func NewCreateOrderBiz(repo CreateOrderRepo) *createOrderBiz {
	return &createOrderBiz{
		repo: repo,
	}
}

func (biz *createOrderBiz) CreateOrder(c context.Context, data *orderModel.OrderCreate) error {
	if err := biz.repo.CreateOrderRepo(c, data); err != nil {
		return common.ErrInternal(err)
	}
	return nil
}
