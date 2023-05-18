package newRepo

import (
	"context"
	"go_service_food_organic/common"
	newModel "go_service_food_organic/module/new/model"
	profileModel "go_service_food_organic/module/profile/model"
	"gorm.io/gorm"
)

type DeleteNewStore interface {
	Delete(c context.Context, id int) error
	FindDataWithCondition(c context.Context, cond map[string]interface{}) (*newModel.New, error)
}

type FindProfileByUserId interface {
	FindDataWithConditon(
		c context.Context,
		cond map[string]interface{},
		morekeys ...string) (*profileModel.Profile, error)
}

type deleteNewRepo struct {
	store        DeleteNewStore
	storeProfile FindProfileByUserId
	req          common.Requester
}

func NewDeleteNewRepo(store DeleteNewStore, storeProfile FindProfileByUserId, req common.Requester) *deleteNewRepo {
	return &deleteNewRepo{
		store:        store,
		storeProfile: storeProfile,
		req:          req,
	}
}

func (repo *deleteNewRepo) DeleteNewRepo(c context.Context, id int) error {
	anew, err := repo.store.FindDataWithCondition(c, map[string]interface{}{"id": id})
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return common.ErrRecordNotFound(newModel.EntityName, err)
		}
		return common.ErrEntityNotExists(newModel.EntityName, err)
	}
	if anew == nil {
		return common.ErrEntityNotExists(newModel.EntityName, nil)
	}
	if anew.Status == 0 {
		return common.ErrEntityDeleted(newModel.EntityName, nil)
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

		if profile.Id != anew.Author {
			return profileModel.ErrorProfileIdNotSame(nil)
		}

		if err := repo.store.Delete(c, anew.Id); err != nil {
			return common.ErrCannotCRUDEntity(newModel.EntityName, common.Delete, err)
		}

		return nil
	}

	if err := repo.store.Delete(c, anew.Id); err != nil {
		return common.ErrCannotCRUDEntity(newModel.EntityName, common.Delete, err)
	}
	return nil
}
