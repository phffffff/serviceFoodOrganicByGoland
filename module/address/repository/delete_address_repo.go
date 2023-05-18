package addressRepo

import (
	"context"
	"go_service_food_organic/common"
	addressModel "go_service_food_organic/module/address/model"
	brandModel "go_service_food_organic/module/brand/model"
	profileModel "go_service_food_organic/module/profile/model"
	"gorm.io/gorm"
)

type DeleteAddressStore interface {
	Delete(c context.Context, id int) error
	FindDataWithCondition(c context.Context, cond map[string]interface{}) (*addressModel.Address, error)
}

type FindProfileByUserIdStore interface {
	FindDataWithConditon(c context.Context, cond map[string]interface{}, morekeys ...string) (*profileModel.Profile, error)
}

type deleteAddressRepo struct {
	store        DeleteAddressStore
	storeProfile FindProfileByUserIdStore
	req          common.Requester
}

func NewDeleteAddressRepo(store DeleteAddressStore, storeProfile FindProfileByUserIdStore, req common.Requester) *deleteAddressRepo {
	return &deleteAddressRepo{
		store:        store,
		storeProfile: storeProfile,
		req:          req,
	}
}

func (repo *deleteAddressRepo) DeleteAddressRepo(c context.Context, id int) error {
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
	if address.Status == 0 {
		return common.ErrEntityDeleted(addressModel.EntityName, nil)
	}

	//kiểm tra xem có phải là địa chỉ của user không
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
		if err := repo.store.Delete(c, address.Id); err != nil {
			return common.ErrCannotCRUDEntity(addressModel.EntityName, common.Delete, err)
		}
		return nil
	}

	if err := repo.store.Delete(c, address.Id); err != nil {
		return common.ErrCannotCRUDEntity(brandModel.EntityName, common.Delete, err)
	}
	return nil
}
