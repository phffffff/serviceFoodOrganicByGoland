package profileRepo

import (
	"context"
	"go_service_food_organic/common"
	profileModel "go_service_food_organic/module/profile/model"
	"reflect"
)

type UpdateProfileStore interface {
	Update(c context.Context, id int, data *profileModel.Profile) error
	FindDataWithConditon(
		c context.Context,
		cond map[string]interface{},
		morekeys ...string,
	) (*profileModel.Profile, error)
}

type updateProfileRepo struct {
	store UpdateProfileStore
	req   common.Requester
}

func NewUpdateProfileRepo(store UpdateProfileStore, req common.Requester) *updateProfileRepo {
	return &updateProfileRepo{store: store, req: req}
}

func (repo *updateProfileRepo) UpdateProfileRepo(
	c context.Context,
	id int,
	data *profileModel.ProfileUpdate,
) error {
	profile, err := repo.store.FindDataWithConditon(c, map[string]interface{}{"id": id})
	if profile == nil && err != nil {
		return common.ErrRecordNotFound(profileModel.EntityName, err)
	}

	if profile.UserId != repo.req.GetUserId() {
		return common.ErrorNoPermission(nil)
	}

	val := reflect.ValueOf(data).Elem()

	for i := 0; i < val.NumField(); i++ {
		field := val.Type().Field(i)
		value := val.Field(i).Interface()

		if value != "" {
			reflect.ValueOf(profile).Elem().FieldByName(field.Name).Set(reflect.ValueOf(value))
		}
	}

	if err := repo.store.Update(c, profile.Id, profile); err != nil {
		return common.ErrInternal(err)
	}
	return nil
}
