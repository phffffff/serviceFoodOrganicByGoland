package foodBusiness

import (
	"context"
	"go_service_food_organic/common"
	foodModel "go_service_food_organic/module/food/model"
)

type UpdateFoodRepo interface {
	UpdateFoodRepo(c context.Context, data *foodModel.FoodUpdate, id int) error
}

type updateFoodBiz struct {
	repo UpdateFoodRepo
}

func NewUpdateFoodBiz(repo UpdateFoodRepo) *updateFoodBiz {
	return &updateFoodBiz{
		repo: repo,
	}
}

func (biz *updateFoodBiz) UpdateFood(c context.Context, data *foodModel.FoodUpdate, id int) error {
	if err := biz.repo.UpdateFoodRepo(c, data, id); err != nil {
		return common.ErrCannotCRUDEntity(foodModel.EntityName, common.Update, err)
	}
	return nil
}
