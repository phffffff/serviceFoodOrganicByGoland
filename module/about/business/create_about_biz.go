package aboutBusiness

import (
	"context"
	"go_service_food_organic/common"
	aboutModel "go_service_food_organic/module/about/model"
)

type CreateAboutStore interface {
	Create(c context.Context, data *aboutModel.About) error
}

type createAboutBiz struct {
	store CreateAboutStore
	req   common.Requester
}

func NewCreateAboutBiz(store CreateAboutStore, req common.Requester) *createAboutBiz {
	return &createAboutBiz{store: store, req: req}
}

func (biz *createAboutBiz) CreateAbout(c context.Context, data *aboutModel.About) error {
	if biz.req.GetRole() != common.Admin {
		return common.ErrorNoPermission(nil)
	}
	if err := biz.store.Create(c, data); err != nil {
		return common.ErrCannotCRUDEntity(aboutModel.EntityName, common.Create, err)
	}
	return nil
}
