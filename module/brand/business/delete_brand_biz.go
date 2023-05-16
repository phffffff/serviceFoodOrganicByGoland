package brandBusiness

import (
	"context"
	"go_service_food_organic/common"
)

type DeleteBrandRepo interface {
	DeleteBrandRepo(c context.Context, id int) error
}

type deleteBrandBiz struct {
	repo DeleteBrandRepo
}

func NewDeleteBrandBiz(repo DeleteBrandRepo) *deleteBrandBiz {
	return &deleteBrandBiz{repo: repo}
}

func (biz *deleteBrandBiz) DeleteBrand(c context.Context, id int) error {
	if err := biz.repo.DeleteBrandRepo(c, id); err != nil {
		return common.ErrInternal(err)
	}
	return nil
}
