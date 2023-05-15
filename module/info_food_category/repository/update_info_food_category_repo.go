package infoFoodCategoryRepo

import (
	"context"
	"go_service_food_organic/common"
	infoFoodCategoryModel "go_service_food_organic/module/info_food_category/model"
	"gorm.io/gorm"
	"reflect"
)

type UpdateInfoFoodCategoryStore interface {
	FindDataWithCondition(c context.Context, cond map[string]interface{}) (*infoFoodCategoryModel.InfoFoodCategory, error)
	Update(c context.Context, id int, data *infoFoodCategoryModel.InfoFoodCategory) error
}

type updateInfoFoodCategoryRepo struct {
	store UpdateInfoFoodCategoryStore
	req   common.Requester
}

func NewUpdateInfoFoodCategoryRepo(store UpdateInfoFoodCategoryStore, req common.Requester) *updateInfoFoodCategoryRepo {
	return &updateInfoFoodCategoryRepo{store: store, req: req}
}

func (repo *updateInfoFoodCategoryRepo) UpdateInfoFoodCategoryRepo(c context.Context, id int, data *infoFoodCategoryModel.InfoFoodCategoryUpdate) error {
	if repo.req.GetRole() != common.Admin {
		return common.ErrorNoPermission(nil)
	}
	ifc, err := repo.store.FindDataWithCondition(c, map[string]interface{}{"id": id})
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return common.ErrRecordNotFound(infoFoodCategoryModel.EntityName, err)
		}
		return common.ErrEntityNotExists(infoFoodCategoryModel.EntityName, err)
	}
	if ifc == nil {
		return common.ErrEntityNotExists(infoFoodCategoryModel.EntityName, nil)
	}

	val := reflect.ValueOf(data).Elem()

	for i := 0; i < val.NumField(); i++ {
		field := val.Type().Field(i)
		value := val.Field(i).Interface()

		if value != "" {
			reflect.ValueOf(ifc).Elem().FieldByName(field.Name).Set(reflect.ValueOf(value))
		}
	}

	if err := repo.store.Update(c, id, ifc); err != nil {
		return common.ErrCannotCRUDEntity(infoFoodCategoryModel.EntityName, common.Update, err)
	}
	return nil
}
