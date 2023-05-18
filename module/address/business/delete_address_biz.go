package addressBusiness

import (
	"context"
	"go_service_food_organic/common"
)

type DeleteAddressRepo interface {
	DeleteAddressRepo(c context.Context, id int) error
}

type deleteAddressBiz struct {
	repo DeleteAddressRepo
}

func NewDeleteAddressBiz(repo DeleteAddressRepo) *deleteAddressBiz {
	return &deleteAddressBiz{repo: repo}
}

func (biz *deleteAddressBiz) DeleteAddress(c context.Context, id int) error {
	if err := biz.repo.DeleteAddressRepo(c, id); err != nil {
		return common.ErrInternal(err)
	}
	return nil
}
