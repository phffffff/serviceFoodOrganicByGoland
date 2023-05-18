package newBusiness

import (
	"context"
	"go_service_food_organic/common"
	newModel "go_service_food_organic/module/new/model"
	profileModel "go_service_food_organic/module/profile/model"
	"gorm.io/gorm"
)

type ListNewStore interface {
	ListDataWithCondition(c context.Context, filter *newModel.Filter, paging *common.Paging) ([]newModel.New, error)
}

type listNewBiz struct {
	store        ListNewStore
	storeProfile FindProfileByUserIdStore
	req          common.Requester
}

func NewListNewBiz(store ListNewStore, storeProfile FindProfileByUserIdStore, req common.Requester) *listNewBiz {
	return &listNewBiz{
		store:        store,
		storeProfile: storeProfile,
		req:          req,
	}
}

func (biz *listNewBiz) ListNew(c context.Context,
	filter *newModel.Filter, paging *common.Paging) ([]newModel.New, error) {

	if biz.req.GetRole() == common.Admin {
		list, err := biz.store.ListDataWithCondition(c, filter, paging)
		if err != nil {
			return nil, common.ErrCannotCRUDEntity(newModel.EntityName, common.List, nil)
		}
		return list, nil
	}
	userId := biz.req.GetUserId()
	if userId == 0 {
		return nil, common.ErrorNoPermission(nil)
	}
	profile, err := biz.storeProfile.FindDataWithConditon(c, map[string]interface{}{"user_id": userId})
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.ErrRecordNotFound(profileModel.EntityName, err)
		}
		return nil, common.ErrEntityNotExists(profileModel.EntityName, err)
	}
	if profile == nil {
		return nil, common.ErrEntityNotExists(profileModel.EntityName, nil)
	}
	if profile.Status == 0 {
		return nil, common.ErrEntityDeleted(profileModel.EntityName, nil)
	}

	filter.Author = profile.Id

	list, err := biz.store.ListDataWithCondition(c, filter, paging)
	if err != nil || list == nil {
		return nil, common.ErrCannotCRUDEntity(newModel.EntityName, common.List, err)
	}
	return list, nil
}
