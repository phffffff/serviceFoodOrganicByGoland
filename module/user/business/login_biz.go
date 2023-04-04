package userBusiness

import (
	"context"
	"go_service_food_organic/common"
	tokenProvider "go_service_food_organic/component/token"
	userModel "go_service_food_organic/module/user/model"
	userRepo "go_service_food_organic/module/user/repository"
)

type LoginStore interface {
	FindDataWithCondition(
		c context.Context,
		data map[string]interface{},
		moreKeys ...string) (*userModel.User, error)
}

type loginBiz struct {
	store         LoginStore
	hasher        userRepo.Hasher
	tokenProvider tokenProvider.Provider
	expiry        int
}

func NewLoginBiz(
	store LoginStore,
	hasher userRepo.Hasher,
	tokenProvider tokenProvider.Provider,
	expiry int,
) *loginBiz {
	return &loginBiz{
		store:         store,
		hasher:        hasher,
		tokenProvider: tokenProvider,
		expiry:        expiry,
	}
}

func (biz *loginBiz) Login(c context.Context, data *userModel.UserLogin) (*tokenProvider.Token, error) {
	user, err := biz.store.FindDataWithCondition(c, map[string]interface{}{"email": data.Email})
	if err != nil && user == nil {
		return nil, userModel.ErrorEmailOrPasswordInvalid(err)
	}

	passhash := biz.hasher.Hash(data.Password + user.Salt)
	if passhash != user.Password {
		return nil, userModel.ErrorEmailOrPasswordInvalid(err)
	}

	payload := tokenProvider.TokenPayload{
		UserId: user.Id,
		Role:   user.Role,
	}

	accessToken, err := biz.tokenProvider.Generate(payload, biz.expiry)
	if err != nil {
		return nil, common.ErrInternal(err)
	}
	return accessToken, nil
}
