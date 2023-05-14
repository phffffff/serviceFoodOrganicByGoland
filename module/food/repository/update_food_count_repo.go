package foodRepo

import (
	"context"
	"go_service_food_organic/common"
	foodModel "go_service_food_organic/module/food/model"
	"gorm.io/gorm"
)

type UpdateFoodCountStore interface {
	UpdateCount(c context.Context, count int, id int, typeOf string) error
	FindDataWithCondition(c context.Context, cond map[string]interface{}) (*foodModel.Food, error)
}
type updateFoodCountRepo struct {
	store UpdateFoodCountStore
}

func NewUpdateFoodCountRepo(store UpdateFoodCountStore) *updateFoodCountRepo {
	return &updateFoodCountRepo{store: store}
}

func (repo *updateFoodCountRepo) UpdateCountFoodRepo(c context.Context, count, id int, typeOf string) error {
	food, err := repo.store.FindDataWithCondition(c, map[string]interface{}{"id": id})
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return common.ErrRecordNotFound(foodModel.EntityName, err)
		}
		return common.ErrEntityNotExists(foodModel.EntityName, err)
	}

	if err := repo.store.UpdateCount(c, count, food.Id, typeOf); err != nil {
		return common.ErrCannotCRUDEntity(foodModel.EntityName, common.Update, err)
	}

	return nil
}
