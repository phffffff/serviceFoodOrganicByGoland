package categoryBusiness

import (
	"context"
	"go_service_food_organic/common"
	categoryModel "go_service_food_organic/module/category/model"
)

type CreateCategoryStore interface {
	Create(c context.Context, data *categoryModel.CategoryCreate) error
}

type createCategoryBiz struct {
	store CreateCategoryStore
	req   common.Requester
}

func NewCreateCategoryBiz(store CreateCategoryStore, req common.Requester) *createCategoryBiz {
	return &createCategoryBiz{store: store, req: req}
}

func (biz *createCategoryBiz) CreateCategory(c context.Context, data *categoryModel.CategoryCreate) error {
	if biz.req.GetRole() != common.Admin {
		return common.ErrorNoPermission(nil)
	}
	if err := biz.store.Create(c, data); err != nil {
		return common.ErrCannotCRUDEntity(categoryModel.EntityName, common.Create, err)
	}
	return nil
}
