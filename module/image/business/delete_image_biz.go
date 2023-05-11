package imageBusiness

import (
	"context"
	"go_service_food_organic/common"
)

type DeleteImageRepo interface {
	DeleteImageRepo(c context.Context, id int) error
}

type deleteImageBiz struct {
	repo DeleteImageRepo
}

func NewDeleteImageBiz(repo DeleteImageRepo) *deleteImageBiz {
	return &deleteImageBiz{repo: repo}
}

func (biz *deleteImageBiz) DeteleImage(c context.Context, id int) error {
	if err := biz.repo.DeleteImageRepo(c, id); err != nil {
		return common.ErrInternal(err)
	}
	return nil
}
