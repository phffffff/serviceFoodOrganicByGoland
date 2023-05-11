package imageFoodModel

import (
	"go_service_food_organic/common"
)

const (
	EntityName          = "Image Food"
	ErrImageTypeInvalid = "ErrImageTypeInvalid"
	MsgImageTypeInvalid = "image type invalid!"
)

type ImageFood struct {
	common.SQLModel `json:",inline"`
	FoodId          int         `json:"-" gorm:"column:food_id"`
	ImageId         int         `json:"-" gorm:"column:image_id"`
	FoodFakeId      *common.UID `json:"food_id" gorm:"-"`
	ImageFakeId     *common.UID `json:"image_id" gorm:"-"`
	Type            string      `json:"type" gorm:"column:type;default:detail;"`
}

func (ImageFood) TableName() string {
	return "image_foods"
}

func (imgf *ImageFood) Mark(isAdminOrOwner bool) {
	imgf.GetUID(common.OjbTypeImage)
	imgf.GetFoodUID(false)
	imgf.GetImageUID(false)
}

func (imgf *ImageFood) GetFoodUID(isAdminOrOwner bool) {
	uid := common.NewUID(uint32(imgf.FoodId), common.OjbTypeFood, 1)
	imgf.FoodFakeId = &uid
}
func (imgf *ImageFood) GetImageUID(isAdminOrOwner bool) {
	uid := common.NewUID(uint32(imgf.FoodId), common.OjbTypeImage, 1)
	imgf.ImageFakeId = &uid
}

type ImageFoodCreate struct {
	common.SQLModel `json:",inline"`
	FoodFakeId      string `json:"food_id" gorm:"-"`
	ImageFakeId     string `json:"image_id" gorm:"-"`
	FoodId          int    `json:"-" gorm:"column:food_id"`
	ImageId         int    `json:"-" gorm:"column:image_id"`
	Type            string `json:"type" gorm:"column:type;default:detail;"`
}

func (ImageFoodCreate) TableName() string {
	return ImageFood{}.TableName()
}

type ImageFoodInfo struct {
	Type    string `json:"type" gorm:"column:type;default:detail;"`
	FoodId  int    `json:"food_id" form:"column:food_id"`
	ImageId int    `json:"image_id" form:"column:image_id"`
}

type ErrorInfo struct {
	FileName string
	ImgInfo  *ImageFoodCreate
	ErrInfo  error
}

func ErrorImageTypeInvalid(err error) *common.AppError {
	return common.NewCustomError(err, MsgImageTypeInvalid, ErrImageTypeInvalid)
}
