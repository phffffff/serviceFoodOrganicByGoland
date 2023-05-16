package brandBusiness

import (
	"context"
	"go_service_food_organic/common"
	brandModel "go_service_food_organic/module/brand/model"
)

type ListBrandStore interface {
	ListDataWithCondition(
		c context.Context,
		filter *brandModel.Filter,
		paging *common.Paging, moreKeys ...string) ([]brandModel.Brand, error)
}

type listBrandBiz struct {
	store ListBrandStore
}

func NewlistBrandBiz(store ListBrandStore) *listBrandBiz {
	return &listBrandBiz{store: store}
}

func (biz *listBrandBiz) ListBrand(c context.Context,
	filter *brandModel.Filter, paging *common.Paging) ([]brandModel.Brand, error) {
	list, err := biz.store.ListDataWithCondition(c, filter, paging, "Foods")
	if err != nil || list == nil {
		return nil, common.ErrCannotCRUDEntity(brandModel.EntityName, common.Create, err)
	}
	return list, nil
}
