package addressBusiness

import (
	"context"
	"go_service_food_organic/common"
	addressModel "go_service_food_organic/module/address/model"
)

type UpdateAddressRepo interface {
	UpdateAddressRepo(c context.Context, id int, data *addressModel.AddressUpdate) error
}

type updateAddressBiz struct {
	repo UpdateAddressRepo
}

func NewUpdateAddressBiz(repo UpdateAddressRepo) *updateAddressBiz {
	return &updateAddressBiz{repo: repo}
}

func (biz *updateAddressBiz) UpdateAddress(c context.Context, id int, data *addressModel.AddressUpdate) error {
	if err := biz.repo.UpdateAddressRepo(c, id, data); err != nil {
		return common.ErrInternal(err)
	}
	return nil
}
