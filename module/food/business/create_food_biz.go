package foodBusiness

import (
	"context"
	"go_service_food_organic/common"
	foodModel "go_service_food_organic/module/food/model"
)

type CreateFoodStore interface {
	Create(c context.Context, data *foodModel.FoodCreate) error
}

type createFoodBiz struct {
	store CreateFoodStore
	req   common.Requester
}

func NewCreateFoodBiz(store CreateFoodStore, req common.Requester) *createFoodBiz {
	return &createFoodBiz{store: store, req: req}
}

func (biz *createFoodBiz) CreateFood(c context.Context, data *foodModel.FoodCreate) error {
	if err := biz.store.Create(c, data); err != nil {
		return common.ErrCannotCRUDEntity(foodModel.EntityName, common.Create, err)
	}
	return nil
}
