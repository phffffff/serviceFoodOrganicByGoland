package newBusiness

import (
	"context"
	"go_service_food_organic/common"
)

type DeleteNewRepo interface {
	DeleteNewRepo(c context.Context, id int) error
}

type deleteNewBiz struct {
	repo DeleteNewRepo
}

func NewDeleteNewBiz(repo DeleteNewRepo) *deleteNewBiz {
	return &deleteNewBiz{repo: repo}
}

func (biz *deleteNewBiz) DeleteNew(c context.Context, id int) error {
	if err := biz.repo.DeleteNewRepo(c, id); err != nil {
		return common.ErrInternal(err)
	}
	return nil
}
