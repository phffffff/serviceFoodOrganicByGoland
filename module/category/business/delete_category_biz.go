package categoryBusiness

import (
	"context"
	"go_service_food_organic/common"
)

type DeleteCategoryRepo interface {
	DeleteCategoryRepo(c context.Context, id int) error
}

type deleteCategoryBiz struct {
	repo DeleteCategoryRepo
}

func NewDeleteCategoryBiz(repo DeleteCategoryRepo) *deleteCategoryBiz {
	return &deleteCategoryBiz{repo: repo}
}

func (biz *deleteCategoryBiz) DeleteCategory(c context.Context, id int) error {
	if err := biz.repo.DeleteCategoryRepo(c, id); err != nil {
		return common.ErrInternal(err)
	}
	return nil
}
