package categoryRepo

import (
	"context"
	"go_service_food_organic/common"
	categoryModel "go_service_food_organic/module/category/model"
	"gorm.io/gorm"
	"reflect"
)

type UpdateCategoryStore interface {
	FindDataWithCondition(c context.Context, cond map[string]interface{}) (*categoryModel.Category, error)
	Update(c context.Context, data *categoryModel.Category, id int) error
}

type updateCategoryRepo struct {
	store UpdateCategoryStore
	req   common.Requester
}

func NewUpdateCategoryRepo(store UpdateCategoryStore, req common.Requester) *updateCategoryRepo {
	return &updateCategoryRepo{
		store: store,
		req:   req,
	}
}

func (repo *updateCategoryRepo) UpdateCategoryRepo(c context.Context, id int, data *categoryModel.CategoryUpdate) error {
	if repo.req.GetRole() != common.Admin {
		return common.ErrorNoPermission(nil)
	}

	category, err := repo.store.FindDataWithCondition(c, map[string]interface{}{"id": id})
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return common.ErrRecordNotFound(categoryModel.EntityName, err)
		}
		return common.ErrEntityNotExists(categoryModel.EntityName, err)
	}
	if category == nil {
		return common.ErrEntityNotExists(categoryModel.EntityName, nil)
	}

	val := reflect.ValueOf(data).Elem()

	for i := 0; i < val.NumField(); i++ {
		field := val.Type().Field(i)
		value := val.Field(i).Interface()

		if value != "" {
			reflect.ValueOf(category).Elem().FieldByName(field.Name).Set(reflect.ValueOf(value))
		}
	}
	if err := repo.store.Update(c, category, category.Id); err != nil {
		return common.ErrCannotCRUDEntity(categoryModel.EntityName, common.Update, err)
	}

	return nil
}
