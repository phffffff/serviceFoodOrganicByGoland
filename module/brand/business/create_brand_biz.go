package brandBusiness

import (
	"context"
	"go_service_food_organic/common"
	brandModel "go_service_food_organic/module/brand/model"
)

type CreateBrandStore interface {
	Create(c context.Context, data *brandModel.BrandCreate) error
}

type createBrandBiz struct {
	store CreateBrandStore
	req   common.Requester
}

func NewCreateBrandBiz(store CreateBrandStore, req common.Requester) *createBrandBiz {
	return &createBrandBiz{store: store, req: req}
}

func (biz *createBrandBiz) CreateBrand(c context.Context, data *brandModel.BrandCreate) error {
	if biz.req.GetRole() != common.Admin {
		return common.ErrorNoPermission(nil)
	}
	if err := biz.store.Create(c, data); err != nil {
		return common.ErrCannotCRUDEntity(brandModel.EntityName, common.Create, err)
	}
	return nil
}
