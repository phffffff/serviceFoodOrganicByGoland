package imageRepo

import (
	"context"
	"go_service_food_organic/common"
	"go_service_food_organic/module/image/model"
)

type DeleteImageStore interface {
	FindDataWithCondition(c context.Context, cond map[string]interface{}) (*imageModel.Image, error)
	Delete(c context.Context, id int) error
}

type deleteImageRepo struct {
	store DeleteImageStore
	req   common.Requester
}

func NewDeleteImageRepo(store DeleteImageStore, req common.Requester) *deleteImageRepo {
	return &deleteImageRepo{store: store, req: req}
}

func (repo *deleteImageRepo) DeleteImageRepo(c context.Context, id int) error {
	if repo.req.GetRole() != common.Admin {
		return common.ErrorNoPermission(nil)
	}
	img, err := repo.store.FindDataWithCondition(c, map[string]interface{}{"id": id})
	if err != nil {
		return common.ErrEntityNotExists(imageModel.EntityName, err)
	}
	if img.Status == 0 {
		return common.ErrEntityDeleted(imageModel.EntityName, nil)
	}
	if err := repo.store.Delete(c, img.Id); err != nil {
		return common.ErrCannotCRUDEntity(imageModel.EntityName, common.Delete, err)
	}
	return nil
}
