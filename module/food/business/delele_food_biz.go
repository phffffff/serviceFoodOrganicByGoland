package foodBusiness

import (
	"context"
	"go_service_food_organic/common"
)

type DeleteFoodRepo interface {
	DeleteFoodRepo(c context.Context, id int) error
}

type deleteFoodBiz struct {
	repo DeleteFoodRepo
}

func NewDeleteFoodBiz(repo DeleteFoodRepo) *deleteFoodBiz {
	return &deleteFoodBiz{repo: repo}
}

func (biz *deleteFoodBiz) DeleteFood(c context.Context, id int) error {
	if err := biz.repo.DeleteFoodRepo(c, id); err != nil {
		return common.ErrInternal(err)
	}
	return nil
}
