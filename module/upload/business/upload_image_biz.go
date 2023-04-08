package uploadBusiness

import (
	"context"
	"fmt"
	"go_service_food_organic/common"
	uploadProvider "go_service_food_organic/component/upload_provider"
	uploadModel "go_service_food_organic/module/upload/model"
	"image"
	"io"
	"log"
	"path/filepath"
	"strings"
	"time"
)

type CreateImageStorage interface {
	Create(c context.Context, data *common.Image) error
}

type uploadImageBiz struct {
	provider uploadProvider.UploadProvider
	store    CreateImageStorage
}

func NewUploadImageBiz(provider uploadProvider.UploadProvider, store CreateImageStorage) *uploadImageBiz {
	return &uploadImageBiz{provider: provider, store: store}
}

func (biz *uploadImageBiz) Upload(c context.Context, data []byte, folder, fileName string) (*common.Image, error) {
	//fileBytes := bytes.NewBuffer(data)
	//
	//w, h, err := getImageDimension(fileBytes)
	//
	//if err != nil {
	//	return nil, uploadModel.ErrFileIsNotImage(err)
	//}
	if strings.TrimSpace(folder) == "" {
		folder = "img"
	}

	fileExt := filepath.Ext(fileName)
	fileName = fmt.Sprintf("%d.%s", time.Now().Nanosecond(), fileExt)

	img, err := biz.provider.SaveFileUploaded(c, data, fmt.Sprintf("%s/%s", folder, fileName))

	if err != nil {
		return nil, uploadModel.CanNotServerSave(err)
	}
	//img.Width = w
	//img.Height = h
	img.Extension = fileExt
	return img, nil
}

func getImageDimension(reader io.Reader) (int, int, error) {
	img, _, err := image.DecodeConfig(reader)
	if err != nil {
		log.Println("err:", err)
		return 0, 0, err
	}
	return img.Width, img.Height, nil
}
