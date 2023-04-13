package imageRepo

import (
	"bytes"
	"context"
	"fmt"
	"go_service_food_organic/common"
	uploadProvider "go_service_food_organic/component/upload_provider"
	imageModel "go_service_food_organic/module/upload/image/model"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"path/filepath"
	"strings"
)

type CreateImageStore interface {
	Create(c context.Context, data *imageModel.Image) error
	FindDataWithCondition(c context.Context, cond map[string]interface{}) (*imageModel.Image, error)
	ListDataWithFilter(
		c context.Context,
		filter *imageModel.Filter,
		paging *common.Paging) ([]imageModel.Image, error)
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
}

func NewUploadImageRepo(
	provider uploadProvider.UploadProvider,
	store CreateImageStore,
	hasher Hasher,
) *uploadImageRepo {
	return &uploadImageRepo{provider: provider, store: store, hasher: hasher}
}

func (repo *uploadImageRepo) UploadImageRepo(c context.Context, data []byte, folder, fileName string) (*imageModel.Image, error) {
	//fileBytes := bytes.NewBuffer(data)

	w, h, err := getImageDimension(data)

	if err != nil {
		return nil, imageModel.ErrFileIsNotImage(err)
	}
	if strings.TrimSpace(folder) == "" {
		folder = "img"
	}

	fileExt := filepath.Ext(fileName)
	hashName := repo.hasher.HashSliceByte(data)
	fileName = fmt.Sprintf("%s%s", hashName, fileExt)

	img, err := repo.provider.SaveFileUploaded(c, data, fmt.Sprintf("%s/%s", folder, fileName))

	if err != nil {
		return nil, imageModel.CanNotServerSave(err)
	}
	img.Width = w
	img.Height = h
	img.Extension = fileExt
	img.HashName = hashName

	//db, _ := repo.store.BeginTransaction()
	//defer func() {
	//	if r := recover(); r != nil {
	//		err := repo.store.RollbackTransaction(db)
	//		if err != nil {
	//			panic(common.ErrDB(err))
	//		}
	//	}
	//}()

	var filter imageModel.Filter
	var paging common.Paging

	paging.FullFill()
	filter.Status = []int{0, 1}

	list, err := repo.store.ListDataWithFilter(c, &filter, &paging)
	if list == nil && err != nil {
		//if err := repo.store.RollbackTransaction(db); err != nil {
		//	return nil, common.ErrorCannotRollback(err)
		//}
		return nil, common.ErrCannotCRUDEntity(imageModel.EntityName, common.List, err)
	}
	for idx := range list {
		if list[idx].HashName == hashName {
			return nil, imageModel.ErrorFileExists()
		}
	}
	if err := repo.store.Create(c, img); err != nil {
		//if err := repo.store.RollbackTransaction(db); err != nil {
		//	return nil, common.ErrorCannotRollback(err)
		//}
		return nil, common.ErrCannotCRUDEntity(imageModel.EntityName, common.Create, err)
	}

	return img, nil
}

func getImageDimension(data []byte) (int, int, error) {
	buffer := bytes.NewBuffer(data)
	img, _, err := image.DecodeConfig(buffer)
	if err != nil {
		return 0, 0, imageModel.ErrorInvalidImageFormat(err)
	}
	return img.Width, img.Height, nil
}
