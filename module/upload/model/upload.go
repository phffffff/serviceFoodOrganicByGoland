package uploadModel

import (
	"errors"
	"go_service_food_organic/common"
)

const (
	ErrFileUploadTooLarge = "ErrFileTooLarge"
	MsgFileTooLarge       = "file too large"

	ErrFileUploadIsNotImage = "ErrFileIsNotImage"
	MsgFileIsNotImage       = "file is not image"

	ErrCanNotSaveFile = "ErrCanNotSaveFile"
	MsgCanNotSaveFile = "can not save file"
)

type Image struct {
	common.SQLModel `json:",inline"`
	FileName        string `json:"file_name" gorm:"column:file_name"'`
	Width           int    `json:"width" gorm:"column:width"`
	Height          int    `json:"height" gorm:"column:height"`
}

func (Image) TableName() string {
	return "images"
}

func ErrFileTooLarge() *common.AppError {
	return common.NewCustomError(
		errors.New(MsgFileTooLarge),
		MsgFileTooLarge,
		ErrFileUploadTooLarge,
	)
}

func ErrFileIsNotImage(err error) *common.AppError {
	return common.NewCustomError(err, MsgFileIsNotImage, ErrFileUploadIsNotImage)
}

func CanNotServerSave(err error) *common.AppError {
	return common.NewCustomError(err, MsgCanNotSaveFile, ErrCanNotSaveFile)
}
