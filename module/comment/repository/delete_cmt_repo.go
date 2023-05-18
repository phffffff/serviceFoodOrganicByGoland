package commentRepo

import (
	"context"
	"go_service_food_organic/common"
	commentModel "go_service_food_organic/module/comment/model"
	profileModel "go_service_food_organic/module/profile/model"
	"gorm.io/gorm"
)

type DeleteCmtStore interface {
	Delete(c context.Context, id int) error
	FindDataWithCondition(c context.Context, cond map[string]interface{}) (*commentModel.Comment, error)
}

type FindProfileByUserId interface {
	FindDataWithConditon(
		c context.Context,
		cond map[string]interface{},
		morekeys ...string) (*profileModel.Profile, error)
}

type deleteCmtRepo struct {
	store        DeleteCmtStore
	storeProfile FindProfileByUserId
	req          common.Requester
}

func NewDeleteCmtRepo(store DeleteCmtStore, storeProfile FindProfileByUserId, req common.Requester) *deleteCmtRepo {
	return &deleteCmtRepo{
		store:        store,
		storeProfile: storeProfile,
		req:          req,
	}
}

func (repo *deleteCmtRepo) DeleteCmtRepo(c context.Context, id int) error {
	cmt, err := repo.store.FindDataWithCondition(c, map[string]interface{}{"id": id})
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return common.ErrRecordNotFound(commentModel.EntityName, err)
		}
		return common.ErrEntityNotExists(commentModel.EntityName, err)
	}
	if cmt == nil {
		return common.ErrEntityNotExists(commentModel.EntityName, nil)
	}
	if cmt.Status == 0 {
		return common.ErrEntityDeleted(commentModel.EntityName, nil)
	}

	//user
	if repo.req.GetRole() != common.Admin {
		profile, err := repo.storeProfile.FindDataWithConditon(c, map[string]interface{}{"user_id": repo.req.GetUserId()})
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

		if profile.Id != cmt.ProfileId {
			return profileModel.ErrorProfileIdNotSame(nil)
		}

		if err := repo.store.Delete(c, cmt.Id); err != nil {
			return common.ErrCannotCRUDEntity(commentModel.EntityName, common.Delete, err)
		}

		return nil
	}

	if err := repo.store.Delete(c, cmt.Id); err != nil {
		return common.ErrCannotCRUDEntity(commentModel.EntityName, common.Delete, err)
	}
	return nil
}
