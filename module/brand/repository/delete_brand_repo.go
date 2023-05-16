package brandRepo

import (
	"context"
	"go_service_food_organic/common"
	brandModel "go_service_food_organic/module/brand/model"
	"gorm.io/gorm"
)

type DeleteBrandStore interface {
	Delete(c context.Context, id int) error
	FindDataWithCondition(c context.Context, cond map[string]interface{}) (*brandModel.Brand, error)
}

type deleteBrandRepo struct {
	store DeleteBrandStore
	req   common.Requester
}

func NewDeleteBrandRepo(store DeleteBrandStore, req common.Requester) *deleteBrandRepo {
	return &deleteBrandRepo{
		store: store,
		req:   req,
	}
}

func (repo *deleteBrandRepo) DeleteBrandRepo(c context.Context, id int) error {
	if repo.req.GetRole() != common.Admin {
		return common.ErrorNoPermission(nil)
	}

	brand, err := repo.store.FindDataWithCondition(c, map[string]interface{}{"id": id})
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return common.ErrRecordNotFound(brandModel.EntityName, err)
		}
		return common.ErrEntityNotExists(brandModel.EntityName, err)
	}
	if brand == nil {
		return common.ErrEntityNotExists(brandModel.EntityName, nil)
	}

	if brand.Status == 0 {
		return common.ErrEntityDeleted(brandModel.EntityName, nil)
	}

	if err := repo.store.Delete(c, brand.Id); err != nil {
		return common.ErrCannotCRUDEntity(brandModel.EntityName, common.Delete, err)
	}
	return nil
}
