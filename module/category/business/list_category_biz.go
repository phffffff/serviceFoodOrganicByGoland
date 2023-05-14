package categoryBusiness

import (
	"context"
	"go_service_food_organic/common"
	categoryModel "go_service_food_organic/module/category/model"
)

type ListCategoryStore interface {
	ListDataWithCondition(
		c context.Context,
		filter *categoryModel.Filter,
		paging *common.Paging, moreKeys ...string) ([]categoryModel.Category, error)
}

type listCategoryBiz struct {
	store ListCategoryStore
}

func NewlistCategoryBiz(store ListCategoryStore) *listCategoryBiz {
	return &listCategoryBiz{store: store}
}

func (biz *listCategoryBiz) ListCategory(c context.Context,
	filter *categoryModel.Filter, paging *common.Paging) ([]categoryModel.Category, error) {
	list, err := biz.store.ListDataWithCondition(c, filter, paging, "Foods")
	if err != nil || list == nil {
		return nil, common.ErrCannotCRUDEntity(categoryModel.EntityName, common.Create, err)
	}
	return list, nil
}
