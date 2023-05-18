package commentBusiness

import (
	"context"
	"go_service_food_organic/common"
)

type DeleteCmtRepo interface {
	DeleteCmtRepo(c context.Context, id int) error
}

type deleteCmtBiz struct {
	repo DeleteCmtRepo
}

func NewDeleteCmtBiz(repo DeleteCmtRepo) *deleteCmtBiz {
	return &deleteCmtBiz{repo: repo}
}

func (biz *deleteCmtBiz) DeleteCmt(c context.Context, id int) error {
	if err := biz.repo.DeleteCmtRepo(c, id); err != nil {
		return common.ErrInternal(err)
	}
	return nil
}
