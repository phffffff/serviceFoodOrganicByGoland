package imageRepo

import (
	"bytes"
	"context"
	"fmt"
	"go_service_food_organic/common"
	uploadProvider "go_service_food_organic/component/upload_provider"
	imageModel2 "go_service_food_organic/module/image/model"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"path/filepath"
	"strings"
)

type CreateImageStore interface {
	Create(c context.Context, data *imageModel2.Image) error
	FindDataWithCondition(c context.Context, cond map[string]interface{}) (*imageModel2.Image, error)
	ListDataWithFilter(
		c context.Context,
		filter *imageModel2.Filter,
		paging *common.Paging) ([]imageModel2.Image, error)
	//BeginTransaction() (*gorm.DB, error)
	//RollbackTransaction(db *gorm.DB) error
	//CommitTransaction(db *gorm.DB) error
}

type Hasher interface {
	HashSliceByte(data []byte) string
}

type uploadImageRepo struct {
	provider uploadProvider.UploadProvider
	store    CreateImageStore
	hasher   Hasher
	req      common.Requester
}

func NewUploadImageRepo(
	provider uploadProvider.UploadProvider,
	store CreateImageStore,
	hasher Hasher,
	req common.Requester,
) *uploadImageRepo {
	return &uploadImageRepo{provider: provider, store: store, hasher: hasher, req: req}
}

func (repo *uploadImageRepo) UploadImageRepo(c context.Context, data []byte, typeImage, fileName string) (*imageModel2.Image, error) {
	//fileBytes := bytes.NewBuffer(data)
	if repo.req.GetRole() != common.Admin {
		return nil, common.ErrorNoPermission(nil)
	}

	w, h, err := getImageDimension(data)

	if err != nil {
		return nil, imageModel2.ErrFileIsNotImage(err)
	}
	if strings.TrimSpace(typeImage) == "" {
		typeImage = "img"
	}

	fileExt := filepath.Ext(fileName)
	hashValue := repo.hasher.HashSliceByte(data)
	fileName = fmt.Sprintf("%s%s", hashValue, fileExt)

	dst := fmt.Sprintf("%s/%s", typeImage, fileName)

	img, err := repo.provider.SaveFileUploaded(c, data, dst)

	if err != nil {
		return nil, imageModel2.CanNotServerSave(err)
	}
	img.Width = w
	img.Height = h
	img.Extension = fileExt
	img.HashValue = hashValue
	img.Type = typeImage

	var filter imageModel2.Filter
	var paging common.Paging

	paging.FullFill()
	filter.Status = []int{0, 1}

	list, err := repo.store.ListDataWithFilter(c, &filter, &paging)
	if list == nil && err != nil {
		//if err := repo.store.RollbackTransaction(db); err != nil {
		//	return nil, common.ErrorCannotRollback(err)
		//}
		return nil, common.ErrCannotCRUDEntity(imageModel2.EntityName, common.List, err)
	}
	for idx := range list {
		if list[idx].HashValue == hashValue && list[idx].Type == typeImage {
			err := imageModel2.ErrorFileExists()
			return nil, err
		}
	}
	if err := repo.store.Create(c, img); err != nil {
		if err := repo.provider.DeleteFileUpload(c, dst); err != nil {
			return nil, imageModel2.CanNotDeleteFileUpload(err)
		}
		return nil, common.ErrCannotCRUDEntity(imageModel2.EntityName, common.Create, err)
	}

	return img, nil
}

func getImageDimension(data []byte) (int, int, error) {
	buffer := bytes.NewBuffer(data)
	img, _, err := image.DecodeConfig(buffer)
	if err != nil {
		return 0, 0, imageModel2.ErrorInvalidImageFormat(err)
	}
	return img.Width, img.Height, nil
}
