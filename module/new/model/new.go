package newModel

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"go_service_food_organic/common"
)

const (
	EntityName = "About"

	StateApproved = "approved"
)

type New struct {
	common.SQLModel `json:",inline"`
	Title           string      `json:"title" gorm:"column:title"`
	Content         string      `json:"content" gorm:"column:content"`
	Image           string      `json:"image" gorm:"column:image"`
	State           string      `json:"state" gorm:"column:state"`
	Author          int         `json:"-" gorm:"column:author"`
	AuthorUID       *common.UID `json:"author" gorm:"-"`
	Tags            *Tags       `json:"tags" gorm:"column:tags"`
}

type Tag struct {
	Name string `json:"name"`
}

func (New) TableName() string {
	return "news"
}

func (new *New) GetAuthorUID(objType int) {
	uid := common.NewUID(uint32(new.Author), objType, 1)
	new.AuthorUID = &uid
}
func (new *New) Mask(isAdminOrOwner bool) {
	new.GetUID(common.OjbTypeNew)
	new.GetAuthorUID(common.OjbTypeProfile)
}

type NewCrt struct {
	common.SQLModel `json:",inline"`
	Title           string `json:"title" gorm:"column:title"`
	Content         string `json:"content" gorm:"column:content"`
	Image           string `json:"image" gorm:"column:image"`
	State           string `json:"state" gorm:"column:state;default:approved"`
	Author          int    `json:"-" gorm:"column:author"`
	AuthorUID       string `json:"author" gorm:"-"`
	Tags            *Tags  `json:"tags" gorm:"column:tags"`
}

func (NewCrt) TableName() string {
	return New{}.TableName()
}
func (n *NewCrt) Mask(isAdminOrOwner bool) {
	n.GetUID(common.OjbTypeNew)
}

type NewUpd struct {
	Title   string `json:"title" gorm:"column:title"`
	Content string `json:"content" gorm:"column:content"`
	Image   string `json:"image" gorm:"column:image"`
	State   string `json:"state" gorm:"column:state;default:approved"`
	Status  int    `json:"status" gorm:"status"`
	Tags    *Tags  `json:"tags" gorm:"column:tags"`
}

func (NewUpd) TableName() string {
	return New{}.TableName()
}

type Tags []Tag

func (tags *Tags) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return common.NewCustomError(nil, fmt.Sprintf("Failed to unmarshal  JSON value: %s", value), "ErrInternal")
	}

	if err := json.Unmarshal(bytes, &tags); err != nil {
		return common.NewCustomError(nil, fmt.Sprintf("Failed to decode  JSON value: %s", value), "ErrInternal")
	}
	return nil
}

func (tags *Tags) Value() (driver.Value, error) {
	if tags == nil {
		return nil, nil
	}
	return json.Marshal(tags)
}
