package addressBusiness

import (
	"context"
	"go_service_food_organic/common"
	addressModel "go_service_food_organic/module/address/model"
	brandModel "go_service_food_organic/module/brand/model"
	profileModel "go_service_food_organic/module/profile/model"
	"gorm.io/gorm"
)

type CreateAddressStore interface {
	Create(c context.Context, data *addressModel.AddressCreate) error
}

type FindProfileByUserIdStore interface {
	FindDataWithConditon(c context.Context, cond map[string]interface{}, morekeys ...string) (*profileModel.Profile, error)
}

type createAddressBiz struct {
	store        CreateAddressStore
	storeProfile FindProfileByUserIdStore
	req          common.Requester
}

func NewCreateAddressBiz(store CreateAddressStore, storeProfile FindProfileByUserIdStore, req common.Requester) *createAddressBiz {
	return &createAddressBiz{store: store, storeProfile: storeProfile, req: req}
}

func (biz *createAddressBiz) CreateAddress(c context.Context, data *addressModel.AddressCreate) error {
	if biz.req.GetRole() == common.Admin {
		if err := biz.store.Create(c, data); err != nil {
			return common.ErrCannotCRUDEntity(addressModel.EntityName, common.Create, err)
		}
		return nil
	}
	userId := biz.req.GetUserId()
	if userId == 0 {
		return common.ErrorNoPermission(nil)
	}
	profile, err := biz.storeProfile.FindDataWithConditon(c, map[string]interface{}{"user_id": userId})
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

	data.ProfileId = profile.Id
	if err := biz.store.Create(c, data); err != nil {
		return common.ErrCannotCRUDEntity(brandModel.EntityName, common.Create, err)
	}
	return nil
}
