package brandRepo

import (
	"context"
	"go_service_food_organic/common"
	brandModel "go_service_food_organic/module/brand/model"
	"gorm.io/gorm"
	"reflect"
)

type UpdateBrandStore interface {
	FindDataWithCondition(c context.Context, cond map[string]interface{}) (*brandModel.Brand, error)
	Update(c context.Context, data *brandModel.Brand, id int) error
}

type updateBrandRepo struct {
	store UpdateBrandStore
	req   common.Requester
}

func NewUpdateBrandRepo(store UpdateBrandStore, req common.Requester) *updateBrandRepo {
	return &updateBrandRepo{
		store: store,
		req:   req,
	}
}

func (repo *updateBrandRepo) UpdateBrandRepo(c context.Context, id int, data *brandModel.BrandUpdate) error {
	if repo.req.GetRole() != common.Admin {
		return common.ErrorNoPermission(nil)
	}

	brand, err := repo.store.FindDataWithCondition(c, map[string]interface{}{"id": id})
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return common.ErrRecordNotFound(brandModel.EntityName, err)
		}
		return common.ErrEntityNotExists(brandModel.EntityName, err)
	}
	if brand == nil {
		return common.ErrEntityNotExists(brandModel.EntityName, nil)
	}

	val := reflect.ValueOf(data).Elem()

	for i := 0; i < val.NumField(); i++ {
		field := val.Type().Field(i)
		value := val.Field(i).Interface()

		if value != "" {
			reflect.ValueOf(brand).Elem().FieldByName(field.Name).Set(reflect.ValueOf(value))
		}
	}
	if err := repo.store.Update(c, brand, brand.Id); err != nil {
		return common.ErrCannotCRUDEntity(brandModel.EntityName, common.Update, err)
	}

	return nil
}
