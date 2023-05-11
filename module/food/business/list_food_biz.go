package foodBusiness

import (
	"context"
	"go_service_food_organic/common"
	foodModel "go_service_food_organic/module/food/model"
)

type ListFoodStore interface {
	ListDataWithCondition(
		c context.Context,
		filter *foodModel.Filter,
		paging *common.Paging,
		moreKeys ...string) ([]foodModel.Food, error)
}

type listFoodBiz struct {
	store ListFoodStore
}

func NewListFoodBiz(store ListFoodStore) *listFoodBiz {
	return &listFoodBiz{store: store}
}

func (biz *listFoodBiz) ListFoodWithFilter(
	c context.Context,
	filter *foodModel.Filter,
	paging *common.Paging) ([]foodModel.Food, error) {

	list, err := biz.store.ListDataWithCondition(c, filter, paging, "FoodImages")
	if err != nil {
		return nil, common.ErrCannotCRUDEntity(foodModel.EntityName, common.List, err)
	}

	return list, nil
}
