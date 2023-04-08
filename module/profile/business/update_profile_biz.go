package profileBusiness

import (
	"context"
	"go_service_food_organic/common"
	profileModel "go_service_food_organic/module/profile/model"
)

type UpdateProfileRepo interface {
	UpdateProfileRepo(c context.Context, id int, data *profileModel.ProfileUpdate) error
}

type updateProfileBiz struct {
	repo UpdateProfileRepo
}

func NewUpdateProfileBiz(repo UpdateProfileRepo) *updateProfileBiz {
	return &updateProfileBiz{repo: repo}
}

func (biz *updateProfileBiz) UpdateProfile(c context.Context, id int, data *profileModel.ProfileUpdate) error {
	if err := biz.repo.UpdateProfileRepo(c, id, data); err != nil {
		return common.ErrCannotCRUDEntity(profileModel.EntityName, common.Update, err)
	}
	return nil
}
