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
	"time"
)

type CreateImageStore interface {
	Create(c context.Context, data *common.Image) error
	FindDataWithCondition(c context.Context, cond map[string]interface{}) (*common.Image, error)
	ListDataWithFilter(
		c context.Context,
		filter *imageModel.Filter,
		paging *common.Paging) ([]common.Image, error)
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

func (repo *uploadImageRepo) UploadImageRepo(c context.Context, data []byte, folder, fileName string) (*common.Image, error) {
	//fileBytes := bytes.NewBuffer(data)

	w, h, err := getImageDimension(data)

	if err != nil {
		return nil, common.ErrFileIsNotImage(err)
	}
	if strings.TrimSpace(folder) == "" {
		folder = "img"
	}

	fileExt := filepath.Ext(fileName)
	fileName = fmt.Sprintf("%d%s", time.Now().Nanosecond(), fileExt)

	img, err := repo.provider.SaveFileUploaded(c, data, fmt.Sprintf("%s/%s", folder, fileName))

	if err != nil {
		return nil, common.CanNotServerSave(err)
	}
	img.Width = w
	img.Height = h
	img.Extension = fileExt
	hashValue := repo.hasher.HashSliceByte(data)
	img.HashValue = hashValue

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
		return nil, common.ErrCannotCRUDEntity(common.EntityName, common.List, err)
	}
	hasFound := false
	for idx := range list {
		if list[idx].HashValue == hashValue {
			hasFound = true
		}
	}
	if !hasFound {
		if err := repo.store.Create(c, img); err != nil {
			//if err := repo.store.RollbackTransaction(db); err != nil {
			//	return nil, common.ErrorCannotRollback(err)
			//}
			return nil, common.ErrCannotCRUDEntity(common.EntityName, common.Create, err)
		}
	}

	return img, nil
}

func getImageDimension(data []byte) (int, int, error) {
	buffer := bytes.NewBuffer(data)
	img, _, err := image.DecodeConfig(buffer)
	if err != nil {
		return 0, 0, common.ErrorInvalidImageFormat(err)
	}
	return img.Width, img.Height, nil
}
