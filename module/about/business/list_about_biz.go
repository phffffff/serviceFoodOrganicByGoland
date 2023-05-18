package aboutBusiness

import (
	"context"
	"go_service_food_organic/common"
	aboutModel "go_service_food_organic/module/about/model"
)

type ListAboutStore interface {
	ListDataWithCondition(
		c context.Context,
		filter *aboutModel.Filter,
		paging *common.Paging) ([]aboutModel.About, error)
}

type listAboutBiz struct {
	store ListAboutStore
}

func NewlistAboutBiz(store ListAboutStore) *listAboutBiz {
	return &listAboutBiz{store: store}
}

func (biz *listAboutBiz) ListAbout(c context.Context,
	filter *aboutModel.Filter, paging *common.Paging) ([]aboutModel.About, error) {
	list, err := biz.store.ListDataWithCondition(c, filter, paging)
	if err != nil || list == nil {
		return nil, common.ErrCannotCRUDEntity(aboutModel.EntityName, common.Create, err)
	}
	return list, nil
}
