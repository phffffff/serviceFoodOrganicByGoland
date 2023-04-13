package uploadProvider

import (
	"context"
	"go_service_food_organic/module/upload/image/model"
)

type UploadProvider interface {
	SaveFileUploaded(c context.Context, data []byte, dst string) (*imageModel.Image, error)
}
