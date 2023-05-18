package aboutBusiness

import (
	"context"
	"go_service_food_organic/common"
)

type DeleteAboutRepo interface {
	DeleteAboutRepo(c context.Context, id int) error
}

type deleteAboutBiz struct {
	repo DeleteAboutRepo
}

func NewDeleteAboutBiz(repo DeleteAboutRepo) *deleteAboutBiz {
	return &deleteAboutBiz{repo: repo}
}

func (biz *deleteAboutBiz) DeleteAbout(c context.Context, id int) error {
	if err := biz.repo.DeleteAboutRepo(c, id); err != nil {
		return common.ErrInternal(err)
	}
	return nil
}
