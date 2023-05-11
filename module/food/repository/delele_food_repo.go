package foodRepo

import (
	"context"
	"go_service_food_organic/common"
	foodModel "go_service_food_organic/module/food/model"
	"gorm.io/gorm"
)

type DeleteFoodStore interface {
	FindDataWithCondition(c context.Context, cond map[string]interface{}) (*foodModel.Food, error)
	Delete(c context.Context, id int) error
}

type deleleFoodRepo struct {
	store DeleteFoodStore
	req   common.Requester
}

func NewDeleteFoodRepo(store DeleteFoodStore, req common.Requester) *deleleFoodRepo {
	return &deleleFoodRepo{
		store: store,
		req:   req,
	}
}

func (repo *deleleFoodRepo) DeleteFoodRepo(c context.Context, id int) error {
	food, err := repo.store.FindDataWithCondition(c, map[string]interface{}{"id": id})
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return common.ErrRecordNotFound(foodModel.EntityName, err)
		}
		return common.ErrEntityNotExists(foodModel.EntityName, err)
	}
	if food.Status == 0 {
		return common.ErrEntityDeleted(foodModel.EntityName, nil)
	}
	if err := repo.store.Delete(c, food.Id); err != nil {
		return common.ErrCannotCRUDEntity(foodModel.EntityName, common.Delete, err)
	}
	return nil
}
