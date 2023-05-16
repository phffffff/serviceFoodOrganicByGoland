package brandBusiness

import (
	"context"
	"go_service_food_organic/common"
	brandModel "go_service_food_organic/module/brand/model"
)

type UpdateBrandRepo interface {
	UpdateBrandRepo(c context.Context, id int, data *brandModel.BrandUpdate) error
}

type updateBrandBiz struct {
	repo UpdateBrandRepo
}

func NewUpdateBrandBiz(repo UpdateBrandRepo) *updateBrandBiz {
	return &updateBrandBiz{repo: repo}
}

func (biz *updateBrandBiz) UpdateBrand(c context.Context, id int, data *brandModel.BrandUpdate) error {
	if err := biz.repo.UpdateBrandRepo(c, id, data); err != nil {
		return common.ErrInternal(err)
	}
	return nil
}
