package userRepo

import (
	"context"
	"go_service_food_organic/common"
	userModel "go_service_food_organic/module/user/model"
	"gorm.io/gorm"
)

type UpdatePassStore interface {
	Update(c context.Context, id int, pass string) error
	FindDataWithCondition(
		c context.Context,
		cond map[string]interface{},
		moreKeys ...string) (*userModel.User, error)
}

type updatePassRepo struct {
	store  UpdatePassStore
	hasher Hasher
	req    common.Requester
}

func NewUpdatePassRepo(store UpdatePassStore, hasher Hasher, req common.Requester) *updatePassRepo {
	return &updatePassRepo{
		store:  store,
		hasher: hasher,
		req:    req,
	}
}

func (repo *updatePassRepo) UpdateUserPassRepo(c context.Context, id int, data *userModel.UserPasswordUpdate) error {
	user, err := repo.store.FindDataWithCondition(c, map[string]interface{}{"id": id})
	if user == nil && err != nil {
		if err == gorm.ErrRecordNotFound {
			return common.ErrRecordNotFound(userModel.EntityName, err)
		}
		return common.ErrEntityNotExists(userModel.EntityName, err)
	}

	//chưa validate

	if user.Id == repo.req.GetUserId() || repo.req.GetRole() == common.Admin {
		salt := user.Salt
		passHash := repo.hasher.Hash(data.Password + salt)
		newPassHash := repo.hasher.Hash(data.NewPassword + salt)

		if passHash == newPassHash {
			return userModel.ErrorNewPassInvalid(nil)
		}

		reNewPassHash := repo.hasher.Hash(data.ReNewPassword + salt)

		if reNewPassHash != newPassHash {
			return userModel.ErrorRePassInvalid(nil)
		}

		if err := repo.store.Update(c, user.Id, newPassHash); err != nil {
			return common.ErrInvalidRequest(err)
		}
	}
	return nil
}
