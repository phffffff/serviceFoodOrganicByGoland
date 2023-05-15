package categoryRepo

import (
	"context"
	"go_service_food_organic/common"
	categoryModel "go_service_food_organic/module/category/model"
	"gorm.io/gorm"
)

type DeleteCategoryStore interface {
	Delete(c context.Context, id int) error
	FindDataWithCondition(c context.Context, cond map[string]interface{}) (*categoryModel.Category, error)
}

type deleteCategoryRepo struct {
	store DeleteCategoryStore
	req   common.Requester
}

func NewDeleteCategoryRepo(store DeleteCategoryStore, req common.Requester) *deleteCategoryRepo {
	return &deleteCategoryRepo{
		store: store,
		req:   req,
	}
}

func (repo *deleteCategoryRepo) DeleteCategoryRepo(c context.Context, id int) error {
	if repo.req.GetRole() != common.Admin {
		return common.ErrorNoPermission(nil)
	}

	category, err := repo.store.FindDataWithCondition(c, map[string]interface{}{"id": id})
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return common.ErrRecordNotFound(categoryModel.EntityName, err)
		}
		return common.ErrEntityNotExists(categoryModel.EntityName, err)
	}
	if category == nil {
		return common.ErrEntityNotExists(categoryModel.EntityName, nil)
	}

	if category.Status == 0 {
		return common.ErrEntityDeleted(categoryModel.EntityName, nil)
	}

	if err := repo.store.Delete(c, category.Id); err != nil {
		return common.ErrCannotCRUDEntity(categoryModel.EntityName, common.Delete, err)
	}
	return nil
}
