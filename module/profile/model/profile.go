package profileModel

import (
	"go_service_food_organic/common"
	"go_service_food_organic/module/image/model"
)

const EntityName = "Profile"

type ProfileRegister struct {
	common.SQLModel
	UserId int    `json:"user_id" gorm:"column:user_id;"`
	Email  string `json:"email" gorm:"column:email;"`
	FbId   string `json:"fb_id" gorm:"column:fb_id;"`
	GgId   string `json:"gg_id" gorm:"column:gg_id;"`
	Phone  string `json:"salt" gorm:"column:phone;"`
}

func (ProfileRegister) TableName() string { return "profiles" }

func (p *ProfileRegister) Mark(isAdminOrOwner bool) {
	p.GetUID(common.OjbTypeProfile)
}

type Profile struct {
	common.SQLModel
	Email        string                   `json:"email" gorm:"column:email;"`
	Phone        string                   `json:"phone" gorm:"column:phone;"`
	LastName     string                   `json:"last_name" gorm:"column:last_name;"`
	FirstName    string                   `json:"first_name" gorm:"column:first_name;"`
	UserId       int                      `json:"-" gorm:"column:user_id;"`
	AvatarId     int                      `json:"-" gorm:"column:avatar_id;"`
	AvatarFakeId *common.UID              `json:"avatar_id" gorm:"-"`
	Image        *imageModel.ImageProfile `json:"image" gorm:"preload:false;foreignKey:AvatarId;"`
}

func (Profile) TableName() string { return ProfileRegister{}.TableName() }

func (p *Profile) Mark(isAdminOrOwner bool) {
	p.GetUID(common.OjbTypeProfile)
	p.GetAvatarUID(common.OjbTypeImage)
}

func (p *Profile) GetAvatarUID(OjbType int) {
	uid := common.NewUID(uint32(p.AvatarId), OjbType, 1)
	p.AvatarFakeId = &uid
}

type ProfileUpdate struct {
	Email        string `json:"email" gorm:"column:email;"`
	Phone        string `json:"phone" gorm:"column:phone;"`
	LastName     string `json:"last_name" gorm:"column:last_name;"`
	FirstName    string `json:"first_name" gorm:"column:first_name;"`
	AvatarId     int    `json:"" gorm:"column:avatar_id;"`
	AvatarFakeId string `json:"avatar_id" gorm:"-"`
}

func (ProfileUpdate) TableName() string { return ProfileUpdate{}.TableName() }
