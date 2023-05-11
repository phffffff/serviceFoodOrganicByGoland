package imageBusiness

import (
	"context"
	"go_service_food_organic/module/image/model"
)

type UploadImageRepo interface {
	UploadImageRepo(c context.Context, data []byte, typeImage, fileName string) (*imageModel.Image, error)
}

type uploadImageBiz struct {
	repo UploadImageRepo
}

func NewUploadImageBiz(repo UploadImageRepo) *uploadImageBiz {
	return &uploadImageBiz{repo: repo}
}

func (biz *uploadImageBiz) Upload(c context.Context, data []byte, typeImage, fileName string) (*imageModel.Image, error) {
	img, err := biz.repo.UploadImageRepo(c, data, typeImage, fileName)
	if err != nil {
		return nil, imageModel.CanNotServerSave(err)
	}
	return img, nil
}
