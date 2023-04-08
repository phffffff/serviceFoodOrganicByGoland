package uploadProvider

import (
	"context"
	"go_service_food_organic/common"
)

type UploadProvider interface {
	SaveFileUploaded(c context.Context, data []byte, dst string) (*common.Image, error)
}
