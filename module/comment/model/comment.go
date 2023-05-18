package commentModel

import "go_service_food_organic/common"

const (
	EntityName = "Comment"
)

type Comment struct {
	common.SQLModel `json:",inline"`
	Content         string      `json:"content" gorm:"column:content"`
	ProfileFakeId   *common.UID `json:"profile_id" gorm:"-"`
	ProfileId       int         `json:"-" gorm:"column:profile_id"`
	NewFakeId       *common.UID `json:"new_id" gorm:"-"`
	NewId           int         `json:"-" gorm:"column:new_id"`
}

func (Comment) TableName() string { return "comments" }

func (c *Comment) GetProfileUID(objType int) {
	uid := common.NewUID(uint32(c.ProfileId), objType, 1)
	c.ProfileFakeId = &uid
}

func (c *Comment) GetNewUID(objType int) {
	uid := common.NewUID(uint32(c.NewId), objType, 1)
	c.NewFakeId = &uid
}
func (c *Comment) Mask(isAdminOrOwner bool) {
	c.GetUID(common.OjbTypeComment)
	c.GetProfileUID(common.OjbTypeProfile)
	c.GetNewUID(common.OjbTypeNew)
}

type CommentCrt struct {
	common.SQLModel `json:",inline"`
	Content         string `json:"content" gorm:"column:content"`
	ProfileId       int    `json:"-" gorm:"column:profile_id"`
	NewId           int    `json:"-" gorm:"column:new_id"`
}

func (CommentCrt) TableName() string { return Comment{}.TableName() }

func (c *CommentCrt) Mask(isAdminOrOwner bool) {
	c.GetUID(common.OjbTypeComment)
}

type CommentUpd struct {
	Content string `json:"content" gorm:"column:content"`
	Status  int    `json:"status" gorm:"column:status"`
}

func (CommentUpd) TableName() string { return Comment{}.TableName() }
