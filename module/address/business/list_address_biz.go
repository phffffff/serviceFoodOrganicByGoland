package addressBusiness

import (
	"context"
	"go_service_food_organic/common"
	addressModel "go_service_food_organic/module/address/model"
)

type ListAddressStore interface {
	ListDataWithCondition(
		c context.Context,
		filter *addressModel.Filter,
		paging *common.Paging) ([]addressModel.Address, error)
}

type listAddressBiz struct {
	store ListAddressStore
}

func NewlistAddressBiz(store ListAddressStore) *listAddressBiz {
	return &listAddressBiz{store: store}
}

func (biz *listAddressBiz) ListAddress(c context.Context,
	filter *addressModel.Filter, paging *common.Paging) ([]addressModel.Address, error) {

	list, err := biz.store.ListDataWithCondition(c, filter, paging)
	if err != nil || list == nil {
		return nil, common.ErrCannotCRUDEntity(addressModel.EntityName, common.Create, err)
	}
	return list, nil
}
