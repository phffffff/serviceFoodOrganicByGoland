package profileRepo

import (
	"context"
	"go_service_food_organic/common"
	profileModel "go_service_food_organic/module/profile/model"
)

type UpdateProfileStore interface {
	Update(c context.Context, data *profileModel.Profile) error
	FindDataWithConditon(
		c context.Context,
		cond map[string]interface{},
		morekeys ...string,
	) (*profileModel.Profile, error)
}

type updateProfileRepo struct {
	store UpdateProfileStore
}

func NewUpdateProfileRepo(store UpdateProfileStore) *updateProfileRepo {
	return &updateProfileRepo{store: store}
}

func (repo *updateProfileRepo) UpdateProfileRepo(c context.Context, id int, data *profileModel.Profile) error {
	profile, err := repo.store.FindDataWithConditon(c, map[string]interface{}{"id": id})
	if profile == nil && err != nil {
		return common.ErrRecordNotFound(profileModel.EntityName, err)
	}

	profile = data

	if err := repo.store.Update(c, profile); err != nil {
		return common.ErrInternal(err)
	}
	return nil
}
