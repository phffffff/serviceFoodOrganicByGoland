package common

type Image struct {
	Id        int    `json:"id" gorm:"column:id;"`
	Url       string `json:"url" gorm:"column:url;"`
	Width     int    `json:"width" gorm:"column:width;"`
	Height    int    `json:"height" gorm:"column:height;"`
	CloudName string `json:"cloud_name,omitempty" gorm:"-"`
	Extension string `json:"extension,omitempty" gorm:"-"`
}

func (Image) TableName() string {
	return "images"
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
