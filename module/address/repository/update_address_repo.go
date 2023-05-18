package addressRepo

import (
	"context"
	"go_service_food_organic/common"
	addressModel "go_service_food_organic/module/address/model"
	brandModel "go_service_food_organic/module/brand/model"
	profileModel "go_service_food_organic/module/profile/model"
	"gorm.io/gorm"
	"reflect"
)

type UpdateAddressStore interface {
	FindDataWithCondition(c context.Context, cond map[string]interface{}) (*addressModel.Address, error)

	Update(c context.Context, data *addressModel.Address, id int) error
}

type updateAddressRepo struct {
	store        UpdateAddressStore
	storeProfile FindProfileByUserIdStore
	req          common.Requester
}

func NewUpdateAddressRepo(store UpdateAddressStore, storeProfile FindProfileByUserIdStore, req common.Requester) *updateAddressRepo {
	return &updateAddressRepo{
		store:        store,
		storeProfile: storeProfile,
		req:          req,
	}
}

func (repo *updateAddressRepo) UpdateAddressRepo(c context.Context, id int, data *addressModel.AddressUpdate) error {
	address, err := repo.store.FindDataWithCondition(c, map[string]interface{}{"id": id})
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return common.ErrRecordNotFound(brandModel.EntityName, err)
		}
		return common.ErrEntityNotExists(addressModel.EntityName, err)
	}
	if address == nil {
		return common.ErrEntityNotExists(addressModel.EntityName, nil)
	}

	if repo.req.GetRole() != common.Admin {
		userId := repo.req.GetUserId()
		if userId == 0 {
			return common.ErrorNoPermission(nil)
		}
		profile, err := repo.storeProfile.FindDataWithConditon(c, map[string]interface{}{"user_id": userId})
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

		if address.ProfileId != profile.Id {
			return profileModel.ErrorProfileIdNotSame(nil)
		}
		if address.Status == 0 {
			return common.ErrEntityDeleted(addressModel.EntityName, nil)
		}

		val := reflect.ValueOf(data).Elem()

		for i := 0; i < val.NumField(); i++ {
			field := val.Type().Field(i)
			value := val.Field(i).Interface()

			if value != "" {
				reflect.ValueOf(address).Elem().FieldByName(field.Name).Set(reflect.ValueOf(value))
			}
		}

		address.ProfileId = profile.Id

		if err := repo.store.Update(c, address, address.Id); err != nil {
			return common.ErrCannotCRUDEntity(addressModel.EntityName, common.Delete, err)
		}
		return nil
	}

	val := reflect.ValueOf(data).Elem()

	for i := 0; i < val.NumField(); i++ {
		field := val.Type().Field(i)
		value := val.Field(i).Interface()

		if value != "" {
			reflect.ValueOf(address).Elem().FieldByName(field.Name).Set(reflect.ValueOf(value))
		}
	}

	if err := repo.store.Update(c, address, address.Id); err != nil {
		return common.ErrCannotCRUDEntity(brandModel.EntityName, common.Update, err)
	}

	return nil
}
