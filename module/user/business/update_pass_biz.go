package userBusiness

import (
	"context"
	"go_service_food_organic/common"
	userModel "go_service_food_organic/module/user/model"
)

type UpdatePassRepo interface {
	UpdateUserPassRepo(c context.Context, id int, data *userModel.UserPasswordUpdate) error
}

type updatePassBiz struct {
	repo UpdatePassRepo
}

func NewUpdatePassBiz(repo UpdatePassRepo) *updatePassBiz {
	return &updatePassBiz{repo: repo}
}

func (biz *updatePassBiz) UpdateUserPassword(c context.Context, id int, data *userModel.UserPasswordUpdate) error {
	if err := biz.repo.UpdateUserPassRepo(c, id, data); err != nil {
		return common.ErrCannotCRUDEntity(userModel.EntityName, common.Update, err)
	}
	return nil
}
