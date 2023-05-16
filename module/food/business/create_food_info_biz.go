package foodBusiness

import (
	"context"
	"go_service_food_organic/common"
	foodModel "go_service_food_organic/module/food/model"
)

type CreateFoodAndInfoRepo interface {
	CreateFoodAndInfoRepo(c context.Context, data *foodModel.FoodCreate, CategoryId int) error
}

type createFoodAndInfoBiz struct {
	repo CreateFoodAndInfoRepo
}

func NewCreateFoodAndInfoBiz(repo CreateFoodAndInfoRepo) *createFoodAndInfoBiz {
	return &createFoodAndInfoBiz{repo: repo}
}

func (biz *createFoodAndInfoBiz) CreateFoodAndInfo(c context.Context, data *foodModel.FoodCreate, categoryId int) error {
	if err := biz.repo.CreateFoodAndInfoRepo(c, data, categoryId); err != nil {
		return common.ErrCannotCRUDEntity(foodModel.EntityName, common.Create, err)
	}
	return nil
}
