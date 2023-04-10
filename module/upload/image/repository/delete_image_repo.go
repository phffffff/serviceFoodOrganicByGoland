package imageRepo

import (
	"context"
	"go_service_food_organic/common"
)

type DeleteImageStore interface {
	FindDataWithCondition(c context.Context, cond map[string]interface{}) (*common.Image, error)
	Delete(c context.Context, id int) error
}

type deleteImageRepo struct {
	store DeleteImageStore
}

func NewDeleteImageRepo(store DeleteImageStore) *deleteImageRepo {
	return &deleteImageRepo{store: store}
}

func (repo *deleteImageRepo) DeleteImageRepo(c context.Context, id int) error {
	img, err := repo.store.FindDataWithCondition(c, map[string]interface{}{"id": id})
	if err != nil && img == nil {
		return common.ErrEntityNotExists(common.EntityName, err)
	}
	if img.Status == 0 {
		return common.ErrEntityDeleted(common.EntityName, nil)
	}
	if err := repo.store.Delete(c, img.Id); err != nil {
		return common.ErrCannotCRUDEntity(common.EntityName, common.Delete, err)
	}
	return nil
}
