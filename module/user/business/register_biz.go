package userBusiness

import (
	"context"
	"go_service_food_organic/common"
	userModel "go_service_food_organic/module/user/model"
)

type RegisterStore interface {
	Create(c context.Context, data *userModel.UserRegister) error
	FindDataWithCondition(
		c context.Context,
		cond map[string]interface{},
		moreKeys ...string) (*userModel.User, error)
}

type Hasher interface {
	Hash(data string) string
}

type registerBiz struct {
	store  RegisterStore
	hasher Hasher
}

func NewRegisterBiz(store RegisterStore, hasher Hasher) *registerBiz {
	return &registerBiz{
		store:  store,
		hasher: hasher,
	}
}

func (biz *registerBiz) Register(c context.Context, data *userModel.UserRegister) error {
	user, _ := biz.store.FindDataWithCondition(c, map[string]interface{}{"email": data.Email})
	if user != nil {
		if user.Status == 0 {
			return userModel.ErrorUserExists()
		}
		return userModel.ErrorUserExists()
	}

	salt := common.GetSalt(50)
	data.Password = biz.hasher.Hash(data.Password + salt)
	data.Salt = salt

	if err := biz.store.Create(c, data); err != nil {
		return common.ErrCannotCRUDEntity(userModel.Entity, common.Create, err)
	}
	return nil
}
