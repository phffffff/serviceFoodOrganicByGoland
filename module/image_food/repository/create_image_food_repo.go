package imageFoodRepo

import (
	"context"
	"go_service_food_organic/common"
	foodModel "go_service_food_organic/module/food/model"
	imageModel "go_service_food_organic/module/image/model"
	imageFoodModel "go_service_food_organic/module/image_food/model"
)

type CreateImageFoodStore interface {
	Create(c context.Context, data *imageFoodModel.ImageFoodCreate) error
}

type FindImageStore interface {
	FindDataWithCondition(c context.Context, cond map[string]interface{}) (*imageModel.Image, error)
}

type FindFoodStore interface {
	FindDataWithCondition(c context.Context, cond map[string]interface{}) (*foodModel.Food, error)
}

type createImageFoodRepo struct {
	store      CreateImageFoodStore
	storeImage FindImageStore
	storeFood  FindFoodStore
	req        common.Requester
}

func NewCreateImageFoodRepo(store CreateImageFoodStore, req common.Requester, storeImage FindImageStore, storeFood FindFoodStore) *createImageFoodRepo {
	return &createImageFoodRepo{store: store, req: req, storeImage: storeImage, storeFood: storeFood}
}

func (repo *createImageFoodRepo) CreateImageFoodRepo(c context.Context, data *imageFoodModel.ImageFoodCreate) error {
	if repo.req.GetRole() != common.Admin {
		return common.ErrorNoPermission(nil)
	}

	imageId := data.ImageId

	img, err := repo.storeImage.FindDataWithCondition(c, map[string]interface{}{"id": imageId})
	if err != nil {
		return common.ErrEntityNotExists(imageModel.EntityName, err)
	}
	if img.Status == 0 {
		return common.ErrRecordNotFound(imageModel.EntityName, err)
	}
	if img.Type != "food" {
		return imageFoodModel.ErrorImageTypeInvalid(nil)
	}

	foodId := data.FoodId
	food, err := repo.storeFood.FindDataWithCondition(c, map[string]interface{}{"id": foodId})
	if err != nil {
		return common.ErrEntityNotExists(foodModel.EntityName, err)
	}
	if food.Status == 0 {
		return common.ErrRecordNotFound(foodModel.EntityName, err)
	}

	if err := repo.store.Create(c, data); err != nil {
		return common.ErrCannotCRUDEntity(imageFoodModel.EntityName, common.Create, err)
	}
	return nil
}
