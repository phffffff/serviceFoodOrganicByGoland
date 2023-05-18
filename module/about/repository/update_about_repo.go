package aboutRepo

import (
	"context"
	"go_service_food_organic/common"
	aboutModel "go_service_food_organic/module/about/model"
	"gorm.io/gorm"
	"reflect"
)

type UpdateAboutStore interface {
	FindDataWithCondition(c context.Context, cond map[string]interface{}) (*aboutModel.About, error)
	Update(c context.Context, data *aboutModel.About, id int) error
}

type updateAboutRepo struct {
	store UpdateAboutStore
	req   common.Requester
}

func NewUpdateAboutRepo(store UpdateAboutStore, req common.Requester) *updateAboutRepo {
	return &updateAboutRepo{
		store: store,
		req:   req,
	}
}

func (repo *updateAboutRepo) UpdateAboutRepo(c context.Context, id int, data *aboutModel.AboutUpdate) error {
	if repo.req.GetRole() != common.Admin {
		return common.ErrorNoPermission(nil)
	}

	about, err := repo.store.FindDataWithCondition(c, map[string]interface{}{"id": id})
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return common.ErrRecordNotFound(aboutModel.EntityName, err)
		}
		return common.ErrEntityNotExists(aboutModel.EntityName, err)
	}
	if about == nil {
		return common.ErrEntityNotExists(aboutModel.EntityName, nil)
	}

	val := reflect.ValueOf(data).Elem()

	for i := 0; i < val.NumField(); i++ {
		field := val.Type().Field(i)
		value := val.Field(i).Interface()

		if value != "" {
			reflect.ValueOf(about).Elem().FieldByName(field.Name).Set(reflect.ValueOf(value))
		}
	}
	if err := repo.store.Update(c, about, about.Id); err != nil {
		return common.ErrCannotCRUDEntity(aboutModel.EntityName, common.Update, err)
	}

	return nil
}
