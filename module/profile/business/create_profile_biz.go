package profileBusiness

import (
	"context"
	"go_service_food_organic/common"
	profileModel "go_service_food_organic/module/profile/model"
)

type CreateProfileStore interface {
	Create(c context.Context, data *profileModel.Profile) error
}

type createProfileBiz struct {
	store CreateProfileStore
}

func NewCreateProfileBiz(store CreateProfileStore) *createProfileBiz {
	return &createProfileBiz{store: store}
}

func (biz *createProfileBiz) CreateProfile(c context.Context, data *profileModel.Profile) error {
	if err := biz.store.Create(c, data); err != nil {
		return common.ErrInternal(err)
	}
	return nil
}
