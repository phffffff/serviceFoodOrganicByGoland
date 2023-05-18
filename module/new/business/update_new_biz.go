package newBusiness

import (
	"context"
	"go_service_food_organic/common"
	newModel "go_service_food_organic/module/new/model"
)

type UpdateNewRepo interface {
	UpdateNewRepo(c context.Context, id int, data *newModel.NewUpd) error
}

type updateNewBiz struct {
	repo UpdateNewRepo
}

func NewUpdateNewBiz(repo UpdateNewRepo) *updateNewBiz {
	return &updateNewBiz{repo: repo}
}

func (biz *updateNewBiz) UpdateNew(c context.Context, id int, data *newModel.NewUpd) error {
	if err := biz.repo.UpdateNewRepo(c, id, data); err != nil {
		return common.ErrInternal(err)
	}
	return nil
}
