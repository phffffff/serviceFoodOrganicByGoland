package commentRepo

import (
	"context"
	"go_service_food_organic/common"
	commentModel "go_service_food_organic/module/comment/model"
	profileModel "go_service_food_organic/module/profile/model"
	"gorm.io/gorm"
	"reflect"
)

type UpdateCmtStore interface {
	FindDataWithCondition(c context.Context, cond map[string]interface{}) (*commentModel.Comment, error)
	Update(c context.Context, data *commentModel.Comment, id int) error
}

type updateCmtRepo struct {
	store        UpdateCmtStore
	storeProfile FindProfileByUserId
	req          common.Requester
}

func NewUpdateCmtRepo(store UpdateCmtStore, storeProfile FindProfileByUserId, req common.Requester) *updateCmtRepo {
	return &updateCmtRepo{
		store:        store,
		storeProfile: storeProfile,
		req:          req,
	}
}

func (repo *updateCmtRepo) UpdateCmtRepo(c context.Context, id int, data *commentModel.CommentUpd) error {
	cmt, err := repo.store.FindDataWithCondition(c, map[string]interface{}{"id": id})
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return common.ErrRecordNotFound(commentModel.EntityName, err)
		}
		return common.ErrEntityNotExists(commentModel.EntityName, err)
	}
	if cmt == nil {
		return common.ErrEntityNotExists(commentModel.EntityName, nil)
	}
	if cmt.Status == 0 {
		return common.ErrEntityDeleted(commentModel.EntityName, nil)
	}

	//user
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

	if profile.Id != cmt.ProfileId {
		return profileModel.ErrorProfileIdNotSame(nil)
	}

	val := reflect.ValueOf(data).Elem()

	for i := 0; i < val.NumField(); i++ {
		field := val.Type().Field(i)
		value := val.Field(i).Interface()

		if value != "" {
			reflect.ValueOf(cmt).Elem().FieldByName(field.Name).Set(reflect.ValueOf(value))
		}
	}
	if err := repo.store.Update(c, cmt, cmt.Id); err != nil {
		return common.ErrCannotCRUDEntity(commentModel.EntityName, common.Delete, err)
	}

	return nil
}
