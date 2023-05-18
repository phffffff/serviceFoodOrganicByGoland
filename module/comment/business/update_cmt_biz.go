package commentBusiness

import (
	"context"
	"go_service_food_organic/common"
	commentModel "go_service_food_organic/module/comment/model"
)

type UpdateCmtRepo interface {
	UpdateCmtRepo(c context.Context, id int, data *commentModel.CommentUpd) error
}

type updateCmtBiz struct {
	repo UpdateCmtRepo
}

func NewUpdateCmtBiz(repo UpdateCmtRepo) *updateCmtBiz {
	return &updateCmtBiz{repo: repo}
}

func (biz *updateCmtBiz) UpdateCmt(c context.Context, id int, data *commentModel.CommentUpd) error {
	if err := biz.repo.UpdateCmtRepo(c, id, data); err != nil {
		return common.ErrInternal(err)
	}
	return nil
}
