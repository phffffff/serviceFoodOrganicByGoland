package imageBusiness

import (
	"context"
	"go_service_food_organic/common"
	imageModel "go_service_food_organic/module/upload/image/model"
)

type ListImageStore interface {
	ListDataWithFilter(
		c context.Context,
		filter *imageModel.Filter,
		paging *common.Paging) ([]imageModel.Image, error)
}

type listImageBiz struct {
	store ListImageStore
	req   common.Requester
}

func NewListImageBiz(store ListImageStore, req common.Requester) *listImageBiz {
	return &listImageBiz{store: store, req: req}
}

func (biz *listImageBiz) ListImage(
	c context.Context,
	filter *imageModel.Filter,
	paging *common.Paging) ([]imageModel.Image, error) {
	list, err := biz.store.ListDataWithFilter(c, filter, paging)

	if biz.req.GetRole() != common.Admin {
		return nil, common.ErrorNoPermission(nil)
	}

	if err != nil && list == nil {
		return nil, common.ErrCannotCRUDEntity(imageModel.EntityName, common.Delete, err)
	}
	return list, nil
}
