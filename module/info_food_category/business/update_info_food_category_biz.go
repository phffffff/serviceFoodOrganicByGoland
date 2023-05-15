package infoFoodCategoryBusiness

import (
	"context"
	"go_service_food_organic/common"
	infoFoodCategoryModel "go_service_food_organic/module/info_food_category/model"
)

type UpdateInfoFoodCategoryRepo interface {
	UpdateInfoFoodCategoryRepo(c context.Context, id int, data *infoFoodCategoryModel.InfoFoodCategoryUpdate) error
}

type updateInfoFoodCategoryBiz struct {
	repo UpdateInfoFoodCategoryRepo
}

func NewUpdateInfoFoodCategoryBiz(repo UpdateInfoFoodCategoryRepo) *updateInfoFoodCategoryBiz {
	return &updateInfoFoodCategoryBiz{repo: repo}
}

func (biz *updateInfoFoodCategoryBiz) UpdateInfoFoodCategory(c context.Context, id int, data *infoFoodCategoryModel.InfoFoodCategoryUpdate) error {
	if err := biz.repo.UpdateInfoFoodCategoryRepo(c, id, data); err != nil {
		return common.ErrInternal(err)
	}
	return nil
}
