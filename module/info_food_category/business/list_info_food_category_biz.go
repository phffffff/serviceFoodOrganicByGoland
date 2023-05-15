package infoFoodCategoryBusiness

import (
	"context"
	"go_service_food_organic/common"
	infoFoodCategoryModel "go_service_food_organic/module/info_food_category/model"
)

type ListInfoFoodCategoryStore interface {
	ListDataWithFilter(
		c context.Context,
		filter *infoFoodCategoryModel.Filter,
		paging *common.Paging,
		moreKeys ...string) ([]infoFoodCategoryModel.InfoFoodCategory, error)
}

type listInfoFoodCategoryBiz struct {
	store ListInfoFoodCategoryStore
	req   common.Requester
}

func NewListInfoFoodCategoryBiz(store ListInfoFoodCategoryStore, req common.Requester) *listInfoFoodCategoryBiz {
	return &listInfoFoodCategoryBiz{store: store, req: req}
}

func (biz *listInfoFoodCategoryBiz) ListInfoFoodCategory(c context.Context, filter *infoFoodCategoryModel.Filter, paging *common.Paging) ([]infoFoodCategoryModel.InfoFoodCategory, error) {
	if biz.req.GetRole() != common.Admin {
		return nil, common.ErrorNoPermission(nil)
	}

	list, err := biz.store.ListDataWithFilter(c, filter, paging, "Foods", "Categories")
	if err != nil || list == nil {
		return nil, common.ErrInternal(err)
	}
	return list, nil
}
