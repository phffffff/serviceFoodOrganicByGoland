package imageFoodBusiness

import (
	"context"
	"go_service_food_organic/common"
	"go_service_food_organic/module/image_food/model"
)

type DeleteImageFoodRepo interface {
	DeleteImageFoodRepo(c context.Context, id int) error
}

type deleteImageFoodBiz struct {
	repo DeleteImageFoodRepo
}

func NewDeleteImageFoodBiz(repo DeleteImageFoodRepo) *deleteImageFoodBiz {
	return &deleteImageFoodBiz{repo: repo}
}

func (biz *deleteImageFoodBiz) DeleteImageFood(c context.Context, id int) error {

	if err := biz.repo.DeleteImageFoodRepo(c, id); err != nil {
		return common.ErrCannotCRUDEntity(imageFoodModel.EntityName, common.Delete, err)
	}
	return nil
}
