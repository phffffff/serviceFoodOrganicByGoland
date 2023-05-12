package cartBusiness

import (
	"context"
	"go_service_food_organic/common"
	cartModel "go_service_food_organic/module/carts/model"
)

type CreateCartRepo interface {
	CreateCartRepo(c context.Context, data *cartModel.Cart) error
}

type createCartBiz struct {
	repo CreateCartRepo
}

func NewCreateCartBiz(repo CreateCartRepo) *createCartBiz {
	return &createCartBiz{repo: repo}
}

func (biz *createCartBiz) CreateCart(c context.Context, data *cartModel.Cart) error {
	if err := biz.repo.CreateCartRepo(c, data); err != nil {
		return common.ErrInternal(err)
	}
	return nil
}
