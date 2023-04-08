package userBusiness

import (
	"context"
	"go_service_food_organic/common"
	userModel "go_service_food_organic/module/user/model"
)

type ListUserStore interface {
	ListDataWithFilter(
		c context.Context,
		filter *userModel.Filter,
		paging *common.Paging,
		moreKeys ...string) ([]userModel.User, error)
}

type listUserBiz struct {
	store ListUserStore
	req   common.Requester
}

func NewListUserBiz(store ListUserStore, req common.Requester) *listUserBiz {
	return &listUserBiz{store: store, req: req}
}

func (biz *listUserBiz) ListUserWithFilter(
	c context.Context,
	filter *userModel.Filter,
	paging *common.Paging) ([]userModel.User, error) {

	if biz.req.GetRole() != common.Admin {
		return nil, common.ErrorNoPermission(nil)
	}

	list, err := biz.store.ListDataWithFilter(c, filter, paging)
	if err != nil {
		return nil, common.ErrCannotCRUDEntity(userModel.EntityName, common.List, err)
	}

	return list, nil
}
