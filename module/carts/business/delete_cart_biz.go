package cartBusiness

import (
	"context"
	"go_service_food_organic/common"
)

type DeleteCartRepo interface {
	DeleteCartRepo(c context.Context) error
}

type deleteCartBiz struct {
	repo DeleteCartRepo
}

func NewDeleteFoodBiz(repo DeleteCartRepo) *deleteCartBiz {
	return &deleteCartBiz{repo: repo}
}

func (biz *deleteCartBiz) DeleteFood(c context.Context) error {
	if err := biz.repo.DeleteCartRepo(c); err != nil {
		return common.ErrInternal(err)
	}
	return nil
}
