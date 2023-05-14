package orderBusiness

import (
	"context"
	"go_service_food_organic/common"
)

type UpdateOrderStateRepo interface {
	UpdateOrderStateRepo(c context.Context, id int, state string) error
}

type updateOrderStateBiz struct {
	repo UpdateOrderStateRepo
}

func NewUpdateOrderStateBiz(repo UpdateOrderStateRepo) *updateOrderStateBiz {
	return &updateOrderStateBiz{repo: repo}
}

func (biz *updateOrderStateBiz) UpdateOrderState(c context.Context, id int, state string) error {
	if err := biz.repo.UpdateOrderStateRepo(c, id, state); err != nil {
		return common.ErrInternal(err)
	}
	return nil
}
