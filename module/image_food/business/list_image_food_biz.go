package imageFoodBusiness

import (
	"context"
	"go_service_food_organic/common"
	imageFoodModel "go_service_food_organic/module/image_food/model"
)

type ListImageFoodStore interface {
	ListDataWithFilter(
		c context.Context,
		filter *imageFoodModel.Filter,
		paging *common.Paging,
		moreKeys ...string) ([]imageFoodModel.ImageFood, error)
}

type listImageFoodBiz struct {
	store ListImageFoodStore
	req   common.Requester
}

func NewListImageFoodBiz(store ListImageFoodStore, req common.Requester) *listImageFoodBiz {
	return &listImageFoodBiz{store: store, req: req}
}

func (biz *listImageFoodBiz) ListImageFood(
	c context.Context,
	filter *imageFoodModel.Filter,
	paging *common.Paging) ([]imageFoodModel.ImageFood, error) {
	list, err := biz.store.ListDataWithFilter(c, filter, paging)

	if biz.req.GetRole() != common.Admin {
		return nil, common.ErrorNoPermission(nil)
	}

	if err != nil && list == nil {
		return nil, common.ErrCannotCRUDEntity(imageFoodModel.EntityName, common.Delete, err)
	}
	return list, nil
}
