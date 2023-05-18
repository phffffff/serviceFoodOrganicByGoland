package commentBusiness

import (
	"context"
	"go_service_food_organic/common"
	commentModel "go_service_food_organic/module/comment/model"
	profileModel "go_service_food_organic/module/profile/model"
	"gorm.io/gorm"
)

type ListCmtStore interface {
	ListDataWithCondition(c context.Context, filter *commentModel.Filter, paging *common.Paging) ([]commentModel.Comment, error)
}

type listCmtBiz struct {
	store        ListCmtStore
	storeProfile FindProfileByUserIdStore
	req          common.Requester
}

func NewListCmtBiz(store ListCmtStore, storeProfile FindProfileByUserIdStore, req common.Requester) *listCmtBiz {
	return &listCmtBiz{
		store:        store,
		storeProfile: storeProfile,
		req:          req,
	}
}

func (biz *listCmtBiz) ListCmt(c context.Context,
	filter *commentModel.Filter, paging *common.Paging) ([]commentModel.Comment, error) {

	if biz.req.GetRole() == common.Admin {
		list, err := biz.store.ListDataWithCondition(c, filter, paging)
		if err != nil {
			return nil, common.ErrCannotCRUDEntity(commentModel.EntityName, common.List, nil)
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

	list, err := biz.store.ListDataWithCondition(c, filter, paging)
	if err != nil || list == nil {
		return nil, common.ErrCannotCRUDEntity(commentModel.EntityName, common.List, err)
	}
	return list, nil
}
