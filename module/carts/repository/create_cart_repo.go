package cartRepo

import (
	"context"
	"go_service_food_organic/common"
	cartModel "go_service_food_organic/module/carts/model"
	userModel "go_service_food_organic/module/user/model"
)

type CreateCartStore interface {
	Create(c context.Context, data *cartModel.Cart) error
}

type createCartRepo struct {
	store CreateCartStore
	req   common.Requester
}

func NewCreateCartRepo(store CreateCartStore, req common.Requester) *createCartRepo {
	return &createCartRepo{store: store, req: req}
}

func (repo *createCartRepo) CreateCartRepo(c context.Context, data *cartModel.Cart) error {
	if repo.req.GetUserId() == 0 {
		return userModel.ErrorUserNoLogin(nil)
	}
	data.UserId = repo.req.GetUserId()

	if err := repo.store.Create(c, data); err != nil {
		return common.ErrCannotCRUDEntity(cartModel.EntityName, common.Create, err)
	}
	return nil
}
