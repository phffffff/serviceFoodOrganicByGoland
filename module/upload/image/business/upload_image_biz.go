package imageBusiness

import (
	"context"
	"go_service_food_organic/common"
)

type UploadImageRepo interface {
	UploadImageRepo(c context.Context, data []byte, folder, fileName string) (*common.Image, error)
}

type uploadImageBiz struct {
	repo UploadImageRepo
}

func NewUploadImageBiz(repo UploadImageRepo) *uploadImageBiz {
	return &uploadImageBiz{repo: repo}
}

func (biz *uploadImageBiz) Upload(c context.Context, data []byte, folder, fileName string) (*common.Image, error) {
	img, err := biz.repo.UploadImageRepo(c, data, folder, fileName)
	if err != nil {
		return nil, common.CanNotServerSave(err)
	}
	return img, nil
}
