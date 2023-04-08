package userRepo

import (
	"context"
	"go_service_food_organic/common"
	profileModel "go_service_food_organic/module/profile/model"
	userModel "go_service_food_organic/module/user/model"
)

type DeleteUserStore interface {
	DeleteUser(c context.Context, idUser int) error
	FindDataWithCondition(
		c context.Context,
		cond map[string]interface{},
		moreKeys ...string,
	) (*userModel.User, error)
}

type DeleteProfileStore interface {
	DeleteProfile(c context.Context, idProfile int) error
	FindDataWithConditon(
		c context.Context,
		cond map[string]interface{},
		morekeys ...string) (*profileModel.Profile, error)
}

type deleteUserRepo struct {
	storeUser    DeleteUserStore
	storeProfile DeleteProfileStore
	req          common.Requester
}

func NewDeleteUserRepo(
	storeUser DeleteUserStore,
	storeProfile DeleteProfileStore,
	req common.Requester,
) *deleteUserRepo {
	return &deleteUserRepo{storeUser: storeUser, storeProfile: storeProfile, req: req}
}

func (biz *deleteUserRepo) DeleteUserRepo(c context.Context, id int) error {
	user, err := biz.storeUser.FindDataWithCondition(c, map[string]interface{}{"id": id})

	if user == nil && err != nil {
		return common.ErrEntityNotExists(userModel.EntityName, err)
	} else {
		if user.Id == biz.req.GetUserId() || biz.req.GetRole() == common.Admin {
			if user.Status == 0 {
				return common.ErrEntityDeleted(userModel.EntityName, nil)
			}
			if err := biz.storeUser.DeleteUser(c, user.Id); err != nil {
				return common.ErrCannotCRUDEntity(userModel.EntityName, common.Delete, err)
			}
			//tìm profile dựa vào user_id
			profile, err := biz.storeProfile.FindDataWithConditon(c, map[string]interface{}{"user_id": user.Id})
			if profile == nil && err != nil {
				return common.ErrEntityNotExists(profileModel.EntityName, err)
			} else {
				//xóa profile dựa vào id của profile
				if err := biz.storeProfile.DeleteProfile(c, profile.Id); err != nil {
					return common.ErrCannotCRUDEntity(profileModel.EntityName, common.Delete, err)
				}
				return nil
			}
		}
		return common.ErrorNoPermission(nil)
	}
}
