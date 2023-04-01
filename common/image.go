package common

type Image struct {
	SQLModel
	FileName string `json:"file_name" gorm:"column:file_name;"`
	Height   int    `json:"height" gorm:"column:height;"`
	Width    int    `json:"width" gorm:"column:width;"`
}

func GetTableName() string { return "images" }

type Images []Image
