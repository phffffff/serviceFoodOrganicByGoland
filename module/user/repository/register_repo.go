package userRepo

import (
	"context"
	"go_service_food_organic/common"
	profileModel "go_service_food_organic/module/profile/model"
	userModel "go_service_food_organic/module/user/model"
)

type RegisterStore interface {
	Create(c context.Context, data *userModel.UserRegister) error
	FindDataWithCondition(c context.Context, cond map[string]interface{}, moreKeys ...string) (*userModel.User, error)
}

type CreateProfileStore interface {
	Create(c context.Context, data *profileModel.ProfileRegister) error
}

type Hasher interface {
	Hash(data string) string
}

type registerRepo struct {
	storeUser    RegisterStore
	hasher       Hasher
	storeProfile CreateProfileStore
}

func NewRegisterRepo(
	storeUser RegisterStore,
	hasher Hasher,
	storeProfile CreateProfileStore) *registerRepo {
	return &registerRepo{
		storeUser:    storeUser,
		hasher:       hasher,
		storeProfile: storeProfile,
	}
}

func (repo *registerRepo) RegisterRepo(c context.Context, data *userModel.UserRegister) error {
	user, err := repo.storeUser.FindDataWithCondition(c, map[string]interface{}{"email": data.Email})
	if user != nil {
		if user.Status == 0 {
			return common.ErrEntityDeleted(userModel.EntityName, err)
		}
		return common.ErrEntityExists(userModel.EntityName, err)
	}

	salt := common.GetSalt(50)
	data.Password = repo.hasher.Hash(data.Password + salt)
	data.Salt = salt

	if err := repo.storeUser.Create(c, data); err != nil {
		return common.ErrCannotCRUDEntity(userModel.EntityName, common.Create, err)
	}

	profile := profileModel.ProfileRegister{
		UserId: data.Id,
		Email:  data.Email,
		FbId:   data.FbId,
		GgId:   data.GgId,
		Phone:  data.Phone,
	}

	if err := repo.storeProfile.Create(c, &profile); err != nil {
		return common.ErrCannotCRUDEntity(profileModel.EntityName, common.Create, err)
	}
	return nil
}
