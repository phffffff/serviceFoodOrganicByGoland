package aboutRepo

import (
	"context"
	"go_service_food_organic/common"
	aboutModel "go_service_food_organic/module/about/model"
	"gorm.io/gorm"
)

type DeleteAboutStore interface {
	Delete(c context.Context, id int) error
	FindDataWithCondition(c context.Context, cond map[string]interface{}) (*aboutModel.About, error)
}

type deleteAboutRepo struct {
	store DeleteAboutStore
	req   common.Requester
}

func NewDeleteAboutRepo(store DeleteAboutStore, req common.Requester) *deleteAboutRepo {
	return &deleteAboutRepo{
		store: store,
		req:   req,
	}
}

func (repo *deleteAboutRepo) DeleteAboutRepo(c context.Context, id int) error {
	if repo.req.GetRole() != common.Admin {
		return common.ErrorNoPermission(nil)
	}

	about, err := repo.store.FindDataWithCondition(c, map[string]interface{}{"id": id})
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return common.ErrRecordNotFound(aboutModel.EntityName, err)
		}
		return common.ErrEntityNotExists(aboutModel.EntityName, err)
	}
	if about == nil {
		return common.ErrEntityNotExists(aboutModel.EntityName, nil)
	}

	if about.Status == 0 {
		return common.ErrEntityDeleted(aboutModel.EntityName, nil)
	}

	if err := repo.store.Delete(c, about.Id); err != nil {
		return common.ErrCannotCRUDEntity(aboutModel.EntityName, common.Delete, err)
	}
	return nil
}
