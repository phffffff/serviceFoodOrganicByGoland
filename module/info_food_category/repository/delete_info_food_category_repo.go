package infoFoodCategoryRepo

import (
	"context"
	"go_service_food_organic/common"
	infoFoodCategoryModel "go_service_food_organic/module/info_food_category/model"
	"gorm.io/gorm"
)

type DeleteInfoFoodCategoryStore interface {
	FindDataWithCondition(c context.Context, cond map[string]interface{}) (*infoFoodCategoryModel.InfoFoodCategory, error)
	Delete(c context.Context, id int) error
}

type deleteInfoFoodCategoryRepo struct {
	store DeleteInfoFoodCategoryStore
	req   common.Requester
}

func NewDeleteInfoFoodCategoryRepo(store DeleteInfoFoodCategoryStore, req common.Requester) *deleteInfoFoodCategoryRepo {
	return &deleteInfoFoodCategoryRepo{store: store, req: req}
}

func (repo *deleteInfoFoodCategoryRepo) DeleteInfoFoodCategoryRepo(c context.Context, id int) error {
	if repo.req.GetRole() != common.Admin {
		return common.ErrorNoPermission(nil)
	}
	ifc, err := repo.store.FindDataWithCondition(c, map[string]interface{}{"id": id})
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return common.ErrRecordNotFound(infoFoodCategoryModel.EntityName, err)
		}
		return common.ErrEntityNotExists(infoFoodCategoryModel.EntityName, err)
	}
	if ifc == nil {
		return common.ErrEntityNotExists(infoFoodCategoryModel.EntityName, nil)
	}

	if ifc.Status == 0 {
		return common.ErrEntityDeleted(infoFoodCategoryModel.EntityName, err)
	}

	if err := repo.store.Delete(c, id); err != nil {
		return common.ErrCannotCRUDEntity(infoFoodCategoryModel.EntityName, common.Delete, err)
	}
	return nil
}
