package userBusiness

import (
	"context"
	"go_service_food_organic/common"
)

type DeleteUserRepo interface {
	DeleteUserRepo(c context.Context, id int) error
}

type deleteUserBiz struct {
	store DeleteUserRepo
}

func NewDeleteUserBiz(store DeleteUserRepo) *deleteUserBiz {
	return &deleteUserBiz{store: store}
}

func (biz *deleteUserBiz) DeleteUser(c context.Context, id int) error {
	if err := biz.store.DeleteUserRepo(c, id); err != nil {
		return common.ErrCannotCRUDEntity(common.CurrentUser, common.Delete, err)
	}
	return nil
}
