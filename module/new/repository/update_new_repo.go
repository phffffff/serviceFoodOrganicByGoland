package newRepo

import (
	"context"
	"go_service_food_organic/common"
	newModel "go_service_food_organic/module/new/model"
	profileModel "go_service_food_organic/module/profile/model"
	"gorm.io/gorm"
	"reflect"
)

type UpdateNewStore interface {
	FindDataWithCondition(c context.Context, cond map[string]interface{}) (*newModel.New, error)
	Update(c context.Context, data *newModel.New, id int) error
}

type updateNewRepo struct {
	store        UpdateNewStore
	storeProfile FindProfileByUserId
	req          common.Requester
}

func NewUpdateNewRepo(store UpdateNewStore, storeProfile FindProfileByUserId, req common.Requester) *updateNewRepo {
	return &updateNewRepo{
		store:        store,
		storeProfile: storeProfile,
		req:          req,
	}
}

func (repo *updateNewRepo) UpdateNewRepo(c context.Context, id int, data *newModel.NewUpd) error {
	anew, err := repo.store.FindDataWithCondition(c, map[string]interface{}{"id": id})
	state := anew.State
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return common.ErrRecordNotFound(newModel.EntityName, err)
		}
		return common.ErrEntityNotExists(newModel.EntityName, err)
	}
	if anew == nil {
		return common.ErrEntityNotExists(newModel.EntityName, nil)
	}

	//user
	if repo.req.GetRole() != common.Admin {
		profile, err := repo.storeProfile.FindDataWithConditon(c, map[string]interface{}{"user_id": repo.req.GetUserId()})
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return common.ErrRecordNotFound(profileModel.EntityName, err)
			}
			return common.ErrEntityNotExists(profileModel.EntityName, err)
		}
		if profile == nil {
			return common.ErrEntityNotExists(profileModel.EntityName, nil)
		}
		if profile.Status == 0 {
			return common.ErrEntityDeleted(profileModel.EntityName, nil)
		}

		if profile.Id != anew.Author {
			return profileModel.ErrorProfileIdNotSame(nil)
		}
		if anew.Status == 0 {
			return common.ErrEntityDeleted(newModel.EntityName, nil)
		}

		val := reflect.ValueOf(data).Elem()

		for i := 0; i < val.NumField(); i++ {
			field := val.Type().Field(i)
			value := val.Field(i).Interface()

			if value != "" {
				reflect.ValueOf(anew).Elem().FieldByName(field.Name).Set(reflect.ValueOf(value))
			}
		}
		anew.State = state
		if err := repo.store.Update(c, anew, anew.Id); err != nil {
			return common.ErrCannotCRUDEntity(newModel.EntityName, common.Delete, err)
		}

		return nil
	}

	//admin

	val := reflect.ValueOf(data).Elem()

	for i := 0; i < val.NumField(); i++ {
		field := val.Type().Field(i)
		value := val.Field(i).Interface()

		if value != "" {
			reflect.ValueOf(anew).Elem().FieldByName(field.Name).Set(reflect.ValueOf(value))
		}
	}
	if err := repo.store.Update(c, anew, anew.Id); err != nil {
		return common.ErrCannotCRUDEntity(newModel.EntityName, common.Update, err)
	}

	return nil
}
