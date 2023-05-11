package imageFoodBusiness

import (
	"context"
	"go_service_food_organic/common"
	"go_service_food_organic/module/image_food/model"
)

type CreateImageFoodRepo interface {
	CreateImageFoodRepo(c context.Context, data *imageFoodModel.ImageFoodCreate) error
}

type createImageFoodBiz struct {
	repo CreateImageFoodRepo
}

func NewCreateImageFoodBiz(repo CreateImageFoodRepo) *createImageFoodBiz {
	return &createImageFoodBiz{repo: repo}
}

func (biz *createImageFoodBiz) CreateImageFood(c context.Context, data *imageFoodModel.ImageFoodCreate) error {
	if err := biz.repo.CreateImageFoodRepo(c, data); err != nil {
		return common.ErrCannotCRUDEntity(imageFoodModel.EntityName, common.Create, err)
	}
	return nil
}
