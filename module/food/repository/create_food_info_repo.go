package foodRepo

import (
	"context"
	"go_service_food_organic/common"
	foodModel "go_service_food_organic/module/food/model"
	infoFoodCategoryModel "go_service_food_organic/module/info_food_category/model"
)

type CreateFoodStore interface {
	Create(c context.Context, data *foodModel.FoodCreate) error
	BeginTransaction() error
	RollbackTransaction() error
	CommitTransaction() error
}

type CreateInfoFoodCategoryStore interface {
	Create(c context.Context, data *infoFoodCategoryModel.InfoFoodCategoryCreate) error
	BeginTransaction() error
	RollbackTransaction() error
	CommitTransaction() error
}

type createFoodAndInfo struct {
	storeFood CreateFoodStore
	storeInfo CreateInfoFoodCategoryStore
	req       common.Requester
}

func NewCreateFoodAndInfo(storeFood CreateFoodStore, storeInfo CreateInfoFoodCategoryStore, req common.Requester) *createFoodAndInfo {
	return &createFoodAndInfo{
		storeFood: storeFood,
		storeInfo: storeInfo,
		req:       req,
	}
}

func (repo *createFoodAndInfo) CreateFoodAndInfoRepo(c context.Context, data *foodModel.FoodCreate, CategoryId int) error {
	if repo.req.GetRole() != common.Admin {
		return common.ErrorNoPermission(nil)
	}
	if err := repo.storeFood.BeginTransaction(); err != nil {
		return common.ErrInternal(err)
	}
	if err := repo.storeInfo.BeginTransaction(); err != nil {
		return common.ErrInternal(err)
	}

	if err := repo.storeFood.Create(c, data); err != nil {
		if err := repo.storeFood.RollbackTransaction(); err != nil {
			return common.ErrInternal(err)
		}
		if err := repo.storeInfo.CommitTransaction(); err != nil {
			return common.ErrInternal(err)
		}
	}

	var ifc infoFoodCategoryModel.InfoFoodCategoryCreate

	ifc.FoodId = data.Id
	ifc.CategoryId = CategoryId

	if err := repo.storeInfo.Create(c, &ifc); err != nil {
		if err := repo.storeFood.RollbackTransaction(); err != nil {
			return common.ErrInternal(err)
		}
		if err := repo.storeInfo.RollbackTransaction(); err != nil {
			return common.ErrInternal(err)
		}
	}

	if err := repo.storeFood.CommitTransaction(); err != nil {
		return common.ErrInternal(err)
	}
	if err := repo.storeInfo.CommitTransaction(); err != nil {
		return common.ErrInternal(err)
	}

	return nil
}
