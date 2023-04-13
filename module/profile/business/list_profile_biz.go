package profileBusiness

import (
	"context"
	"go_service_food_organic/common"
	profileModel "go_service_food_organic/module/profile/model"
)

type ListProfileStore interface {
	ListDataWithFilter(
		c context.Context,
		filter *profileModel.Filter,
		paging *common.Paging,
		moreKeys ...string) ([]profileModel.Profile, error)
}

type listProfileBiz struct {
	store ListProfileStore
	req   common.Requester
}

func NewListProfileBiz(store ListProfileStore, req common.Requester) *listProfileBiz {
	return &listProfileBiz{store: store, req: req}
}

func (biz *listProfileBiz) ListProfileWithFilter(
	c context.Context,
	filter *profileModel.Filter,
	paging *common.Paging) ([]profileModel.Profile, error) {

	if biz.req.GetRole() != common.Admin {
		return nil, common.ErrorNoPermission(nil)
	}

	list, err := biz.store.ListDataWithFilter(c, filter, paging, "Image")
	if err != nil {
		return nil, common.ErrCannotCRUDEntity(profileModel.EntityName, common.List, err)
	}

	return list, nil
}
