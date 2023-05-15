package infoFoodCategoryBusiness

import (
	"context"
	"go_service_food_organic/common"
	infoFoodCategoryModel "go_service_food_organic/module/info_food_category/model"
)

type CreateInfoFoodCategoryStore interface {
	Create(c context.Context, data *infoFoodCategoryModel.InfoFoodCategoryCreate) error
}

type createInfoFoodCategoryBiz struct {
	store CreateInfoFoodCategoryStore
	req   common.Requester
}

func NewCreateInfoFoodCategoryBiz(store CreateInfoFoodCategoryStore, req common.Requester) *createInfoFoodCategoryBiz {
	return &createInfoFoodCategoryBiz{store: store, req: req}
}

func (biz *createInfoFoodCategoryBiz) CreateInfoFoodCategory(c context.Context, data *infoFoodCategoryModel.InfoFoodCategoryCreate) error {
	if biz.req.GetRole() != common.Admin {
		return common.ErrorNoPermission(nil)
	}

	if err := biz.store.Create(c, data); err != nil {
		return common.ErrInternal(err)
	}
	return nil
}
