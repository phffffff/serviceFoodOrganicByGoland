package infoFoodCategoryBusiness

import (
	"context"
	"go_service_food_organic/common"
)

type DeleteInfoFoodCategoryRepo interface {
	DeleteInfoFoodCategoryRepo(c context.Context, id int) error
}

type deleteInfoFoodCategoryBiz struct {
	repo DeleteInfoFoodCategoryRepo
}

func NewDeleteInfoFoodCategoryBiz(repo DeleteInfoFoodCategoryRepo) *deleteInfoFoodCategoryBiz {
	return &deleteInfoFoodCategoryBiz{repo: repo}
}

func (biz *deleteInfoFoodCategoryBiz) DeleteInfoFoodCategory(c context.Context, id int) error {
	if err := biz.repo.DeleteInfoFoodCategoryRepo(c, id); err != nil {
		return common.ErrInternal(err)
	}
	return nil
}
