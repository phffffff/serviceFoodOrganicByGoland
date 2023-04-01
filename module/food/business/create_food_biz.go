package foodBusiness

import (
	"context"
	foodModel "go_service_food_organic/module/food/model"
)

type CreateFoodStore interface {
	Create(c context.Context, data *foodModel.FoodCreate) error
}

type createFoodBiz struct {
	store CreateFoodStore
}

func NewCreateFoodBiz(store CreateFoodStore) *createFoodBiz {
	return &createFoodBiz{store: store}
}

func (biz *createFoodBiz) CreateFood(c context.Context, data *foodModel.FoodCreate) error {
	if err := biz.store.Create(c, data); err != nil {
		return err
	}
	return nil
}
