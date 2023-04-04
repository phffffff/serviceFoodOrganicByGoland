package userBusiness

import (
	"context"
	"go_service_food_organic/common"
	userModel "go_service_food_organic/module/user/model"
)

type RegisterRepo interface {
	RegisterRepo(c context.Context, data *userModel.UserRegister) error
}

type registerBiz struct {
	repo RegisterRepo
}

func NewRegisterBiz(repo RegisterRepo) *registerBiz {
	return &registerBiz{repo: repo}
}

func (biz *registerBiz) Register(c context.Context, data *userModel.UserRegister) error {
	if err := biz.repo.RegisterRepo(c, data); err != nil {
		return common.ErrInternal(err)
	}
	return nil
}
