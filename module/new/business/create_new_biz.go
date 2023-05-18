package newBusiness

import (
	"context"
	"go_service_food_organic/common"
	newModel "go_service_food_organic/module/new/model"
	profileModel "go_service_food_organic/module/profile/model"
	"gorm.io/gorm"
)

type CreateNewStore interface {
	Create(c context.Context, data *newModel.NewCrt) error
}

type FindProfileByUserIdStore interface {
	FindDataWithConditon(c context.Context, cond map[string]interface{}, morekeys ...string) (*profileModel.Profile, error)
}

type createNewBiz struct {
	store        CreateNewStore
	storeProfile FindProfileByUserIdStore
	req          common.Requester
}

func NewCreateNewBiz(store CreateNewStore, storeProfile FindProfileByUserIdStore, req common.Requester) *createNewBiz {
	return &createNewBiz{store: store, storeProfile: storeProfile, req: req}
}

func (biz *createNewBiz) CreateNew(c context.Context, data *newModel.NewCrt) error {
	//admin
	if biz.req.GetRole() == common.Admin {
		if err := biz.store.Create(c, data); err != nil {
			return common.ErrCannotCRUDEntity(newModel.EntityName, common.Create, err)
		}
		return nil
	}
	//user
	userId := biz.req.GetUserId()
	if userId == 0 {
		return common.ErrorNoPermission(nil)
	}
	profile, err := biz.storeProfile.FindDataWithConditon(c, map[string]interface{}{"user_id": userId})
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return common.ErrRecordNotFound(profileModel.EntityName, err)
		}
		return common.ErrEntityNotExists(profileModel.EntityName, err)
	}
	if profile == nil {
		return common.ErrEntityNotExists(profileModel.EntityName, nil)
	}
	if profile.Status == 0 {
		return common.ErrEntityDeleted(profileModel.EntityName, nil)
	}

	data.Author = profile.Id
	data.State = newModel.StateApproved

	if err := biz.store.Create(c, data); err != nil {
		return common.ErrCannotCRUDEntity(newModel.EntityName, common.Create, err)
	}
	return nil
}
