package uploadProvider

import (
	"context"
	"go_service_food_organic/module/image/model"
)

type UploadProvider interface {
	SaveFileUploaded(c context.Context, data []byte, dst string) (*imageModel.Image, error)
	DeleteFileUpload(c context.Context, dst string) error
}
