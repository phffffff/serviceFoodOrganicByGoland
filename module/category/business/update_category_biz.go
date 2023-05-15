package categoryBusiness

import (
	"context"
	"go_service_food_organic/common"
	categoryModel "go_service_food_organic/module/category/model"
)

type UpdateCategoryRepo interface {
	UpdateCategoryRepo(c context.Context, id int, data *categoryModel.CategoryUpdate) error
}

type updateCategoryBiz struct {
	repo UpdateCategoryRepo
}

func NewUpdateCategoryBiz(repo UpdateCategoryRepo) *updateCategoryBiz {
	return &updateCategoryBiz{repo: repo}
}

func (biz *updateCategoryBiz) UpdateCategory(c context.Context, id int, data *categoryModel.CategoryUpdate) error {
	if err := biz.repo.UpdateCategoryRepo(c, id, data); err != nil {
		return common.ErrInternal(err)
	}
	return nil
}
