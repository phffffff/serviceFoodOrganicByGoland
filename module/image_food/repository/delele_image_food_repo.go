package imageFoodRepo

import (
	"context"
	"go_service_food_organic/common"
	"go_service_food_organic/module/image_food/model"
	"gorm.io/gorm"
)

type DeleteImageFoodStore interface {
	Delete(c context.Context, id int) error
	FindDataWithCondition(c context.Context, cond map[string]interface{}) (*imageFoodModel.ImageFood, error)
}

type deleteImageFoodRepo struct {
	store DeleteImageFoodStore
	req   common.Requester
}

func NewDeleteImageFoodRepo(store DeleteImageFoodStore, req common.Requester) *deleteImageFoodRepo {
	return &deleteImageFoodRepo{store: store, req: req}
}

func (repo *deleteImageFoodRepo) DeleteImageFoodRepo(c context.Context, id int) error {
	if repo.req.GetRole() != common.Admin {
		return common.ErrorNoPermission(nil)
	}
	imf, err := repo.store.FindDataWithCondition(c, map[string]interface{}{"id": id})
	if imf == nil && err != nil {
		if err == gorm.ErrRecordNotFound {
			return common.ErrEntityNotExists(imageFoodModel.EntityName, err)
		}
		return common.ErrInternal(err)
	}
	if imf.Status == 0 {
		return common.ErrEntityDeleted(imageFoodModel.EntityName, nil)
	}
	if err := repo.store.Delete(c, imf.Id); err != nil {
		return common.ErrCannotCRUDEntity(imageFoodModel.EntityName, common.Delete, err)
	}
	return nil
}
