package common

import "errors"

const (
	EntityName = "Image"

	ErrFileUploadTooLarge = "ErrFileTooLarge"
	MsgFileTooLarge       = "file too large"

	ErrFileUploadIsNotImage = "ErrFileIsNotImage"
	MsgFileIsNotImage       = "file is not image"

	ErrCanNotSaveFile = "ErrCanNotSaveFile"
	MsgCanNotSaveFile = "can not save file"

	ErrInvalidImageFormat = "ErrInvalidImageFormat"
	MsgInvalidImageFormat = "unknown format image"
)

type Image struct {
	SQLModel  `json:",inline"`
	Url       string `json:"url" gorm:"column:url;"`
	Width     int    `json:"width" gorm:"column:width;"`
	Height    int    `json:"height" gorm:"column:height;"`
	HashValue string `json:"hash_value" gorm:"column:hash_value;"`
	CloudName string `json:"cloud_name,omitempty" gorm:"-"`
	Extension string `json:"extension,omitempty" gorm:"-"`
}

func (Image) GetTableName() string {
	return "images"
}

func (img *Image) Mark(isAdminOrOwner bool) {
	img.GetUID(OjbTypeImage)
}

func ErrorInvalidImageFormat(err error) *AppError {
	return NewCustomError(err, MsgInvalidImageFormat, ErrInvalidImageFormat)
}

func ErrFileTooLarge() *AppError {
	return NewCustomError(
		errors.New(MsgFileTooLarge),
		MsgFileTooLarge,
		ErrFileUploadTooLarge,
	)
}

func ErrFileIsNotImage(err error) *AppError {
	return NewCustomError(err, MsgFileIsNotImage, ErrFileUploadIsNotImage)
}

func CanNotServerSave(err error) *AppError {
	return NewCustomError(err, MsgCanNotSaveFile, ErrCanNotSaveFile)
}

//
//func (img *Image) Scan(value interface{}) error {
//	bytes, ok := value.([]byte)
//	if !ok {
//		return NewCustomError(nil, "Failed to unmarshal  JSON value", "ErrInternal")
//	}
//
//	var newImg Image
//	if err := json.Unmarshal(bytes, &newImg); err != nil {
//		return NewCustomError(nil, "Failed to decode  JSON value", "ErrInternal")
//	}
//
//	*img = newImg
//
//	return nil
//}
//
//func (img *Image) Value() (driver.Value, error) {
//	if img == nil {
//		return nil, nil
//	}
//	return json.Marshal(img)
//}
//
//type Images []Image
//
//func (imgs *Images) Scan(value interface{}) error {
//	bytes, ok := value.([]byte)
//	if !ok {
//		return NewCustomError(
//			nil,
//			fmt.Sprintf("Failed to unmarshal  JSON value: %s", value),
//			"ErrInternal")
//	}
//
//	var newImgs Images
//	if err := json.Unmarshal(bytes, &newImgs); err != nil {
//		return NewCustomError(
//			nil,
//			fmt.Sprintf("Failed to decode  JSON value: %s", value),
//			"ErrInternal")
//	}
//	*imgs = newImgs
//
//	return nil
//}
//
//func (imgs *Images) Value() (driver.Value, error) {
//	if imgs == nil {
//		return nil, nil
//	}
//	return json.Marshal(imgs)
//}
