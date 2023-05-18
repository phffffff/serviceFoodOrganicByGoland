package aboutBusiness

import (
	"context"
	"go_service_food_organic/common"
	aboutModel "go_service_food_organic/module/about/model"
)

type UpdateAboutRepo interface {
	UpdateAboutRepo(c context.Context, id int, data *aboutModel.AboutUpdate) error
}

type updateAboutBiz struct {
	repo UpdateAboutRepo
}

func NewUpdateAboutBiz(repo UpdateAboutRepo) *updateAboutBiz {
	return &updateAboutBiz{repo: repo}
}

func (biz *updateAboutBiz) UpdateAbout(c context.Context, id int, data *aboutModel.AboutUpdate) error {
	if err := biz.repo.UpdateAboutRepo(c, id, data); err != nil {
		return common.ErrInternal(err)
	}
	return nil
}
