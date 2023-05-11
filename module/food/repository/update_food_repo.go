package foodRepo

import (
	"context"
	"go_service_food_organic/common"
	foodModel "go_service_food_organic/module/food/model"
	"gorm.io/gorm"
	"reflect"
)

type UpdateFoodStore interface {
	FindDataWithCondition(c context.Context, cond map[string]interface{}) (*foodModel.Food, error)
	Update(c context.Context, data *foodModel.Food, id int) error
}

type updateFoodRepo struct {
	store UpdateFoodStore
	req   common.Requester
}

func NewUpdateFoodRepo(store UpdateFoodStore, req common.Requester) *updateFoodRepo {
	return &updateFoodRepo{store: store, req: req}
}

func (repo *updateFoodRepo) UpdateFoodRepo(c context.Context, data *foodModel.FoodUpdate, id int) error {
	if repo.req.GetRole() != common.Admin {
		return common.ErrorNoPermission(nil)
	}
	food, err := repo.store.FindDataWithCondition(c, map[string]interface{}{"id": id})
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return common.ErrRecordNotFound(foodModel.EntityName, err)
		}
		return common.ErrEntityNotExists(foodModel.EntityName, err)
	}

	val := reflect.ValueOf(data).Elem()

	for i := 0; i < val.NumField(); i++ {
		field := val.Type().Field(i)
		value := val.Field(i).Interface()

		if value != "" {
			reflect.ValueOf(food).Elem().FieldByName(field.Name).Set(reflect.ValueOf(value))
		}
	}

	foodId := food.Id
	if err := repo.store.Update(c, food, foodId); err != nil {
		return common.ErrCannotCRUDEntity(foodModel.EntityName, common.Update, err)
	}
	return nil
}
