package imageFoodRepo

import (
	"context"
	"go_service_food_organic/common"
	imageFoodModel "go_service_food_organic/module/image_food/model"
	"gorm.io/gorm"
)

type UpdateImageFoodStore interface {
	Delete(c context.Context, id int) error
	Create(c context.Context, data *imageFoodModel.ImageFoodCreate) error
	FindDataWithCondition(c context.Context, cond map[string]interface{}) (*imageFoodModel.ImageFood, error)
	BeginTransaction() error
	RollbackTransaction() error
	CommitTransaction() error
}

type updateImageFoodRepo struct {
	store UpdateImageFoodStore
	req   common.Requester
}

func NewUpdateImageFoodRepo(store UpdateImageFoodStore, req common.Requester) *updateImageFoodRepo {
	return &updateImageFoodRepo{
		store: store,
		req:   req,
	}
}

func (repo *updateImageFoodRepo) UpdateImageFoodRepo(c context.Context, data *imageFoodModel.ImageFoodCreate, id int) error {
	if repo.req.GetRole() != common.Admin {
		return common.ErrorNoPermission(nil)
	}

	//begin transaction
	if err := repo.store.BeginTransaction(); err != nil {
		return common.ErrInternal(err)
	}

	imf, err := repo.store.FindDataWithCondition(c, map[string]interface{}{"id": id})
	if imf == nil && err != nil {
		if err := repo.store.CommitTransaction(); err != nil {
			return common.ErrInternal(err)
		}
		if err == gorm.ErrRecordNotFound {
			return common.ErrEntityNotExists(imageFoodModel.EntityName, err)
		}
		return common.ErrInternal(err)
	}

	if imf.Status == 0 {
		if err := repo.store.CommitTransaction(); err != nil {
			return common.ErrInternal(err)
		}
		return common.ErrEntityDeleted(imageFoodModel.EntityName, nil)
	}

	if err := repo.store.Delete(c, imf.Id); err != nil {
		if err := repo.store.RollbackTransaction(); err != nil {
			return common.ErrInternal(err)
		}
		return common.ErrCannotCRUDEntity(imageFoodModel.EntityName, common.Delete, err)
	}
	if err := repo.store.Create(c, data); err != nil {
		if err := repo.store.RollbackTransaction(); err != nil {
			return common.ErrInternal(err)
		}
		return common.ErrCannotCRUDEntity(imageFoodModel.EntityName, common.Create, err)
	}
	if err := repo.store.CommitTransaction(); err != nil {
		return common.ErrInternal(err)
	}
	return nil
}
