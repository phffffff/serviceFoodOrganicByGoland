package cartRepo

import (
	"context"
	"go_service_food_organic/common"
	cartModel "go_service_food_organic/module/carts/model"
	userModel "go_service_food_organic/module/user/model"
)

type DeleteCartStore interface {
	FindDataWithCondition(c context.Context, cond map[string]interface{}) (*cartModel.Cart, error)
	Delete(c context.Context, userId int) error
}

type deleleCartRepo struct {
	store DeleteCartStore
	req   common.Requester
}

func NewDeleteCartRepo(store DeleteCartStore, req common.Requester) *deleleCartRepo {
	return &deleleCartRepo{
		store: store,
		req:   req,
	}
}

func (repo *deleleCartRepo) DeleteCartRepo(c context.Context) error {
	if repo.req.GetUserId() == 0 {
		return userModel.ErrorUserNoLogin(nil)
	}
	userId := repo.req.GetUserId()
	if err := repo.store.Delete(c, userId); err != nil {
		return common.ErrCannotCRUDEntity(cartModel.EntityName, common.Delete, err)
	}
	return nil
}
