package commentBusiness

import (
	"context"
	"go_service_food_organic/common"
	commentModel "go_service_food_organic/module/comment/model"
	profileModel "go_service_food_organic/module/profile/model"
	"gorm.io/gorm"
)

type CreateCmtStore interface {
	Create(c context.Context, data *commentModel.CommentCrt) error
}

type FindProfileByUserIdStore interface {
	FindDataWithConditon(c context.Context, cond map[string]interface{}, morekeys ...string) (*profileModel.Profile, error)
}

type createCmtBiz struct {
	store        CreateCmtStore
	storeProfile FindProfileByUserIdStore
	req          common.Requester
}

func NewCreateCmtBiz(store CreateCmtStore, storeProfile FindProfileByUserIdStore, req common.Requester) *createCmtBiz {
	return &createCmtBiz{store: store, storeProfile: storeProfile, req: req}
}

func (biz *createCmtBiz) CreateCmt(c context.Context, data *commentModel.CommentCrt) error {
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

	data.ProfileId = profile.Id

	if err := biz.store.Create(c, data); err != nil {
		return common.ErrCannotCRUDEntity(commentModel.EntityName, common.Create, err)
	}
	return nil
}
