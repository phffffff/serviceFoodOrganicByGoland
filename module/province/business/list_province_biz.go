package provinceBusiness

import (
	"context"
	"go_service_food_organic/common"
	provinceModel "go_service_food_organic/module/province/model"
)

type ListProvinceStore interface {
	ListDataWithFilter(
		c context.Context,
		filter *provinceModel.Filter,
		paging *common.Paging) ([]provinceModel.Province, error)
}

type listProvinceBiz struct {
	store ListProvinceStore
}

func NewListProvinceBiz(store ListProvinceStore) *listProvinceBiz {
	return &listProvinceBiz{store: store}
}

func (biz *listProvinceBiz) ListProvince(c context.Context, filter *provinceModel.Filter, paging *common.Paging) ([]provinceModel.Province, error) {
	list, err := biz.store.ListDataWithFilter(c, filter, paging)
	if err != nil || list == nil {
		return nil, common.ErrCannotCRUDEntity(provinceModel.EntityName, common.List, err)
	}
	return list, nil
}
