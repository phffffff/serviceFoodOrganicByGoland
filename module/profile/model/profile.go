package profileModel

import (
	"go_service_food_organic/common"
	imageModel "go_service_food_organic/module/upload/image/model"
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

func (ProfileRegister) GetTableName() string { return "profiles" }

func (p *ProfileRegister) Mark(isAdminOrOwner bool) {
	p.GetUID(common.OjbTypeProfile)
}

type Profile struct {
	common.SQLModel
	Email     string                  `json:"email" gorm:"column:email;"`
	Phone     string                  `json:"phone" gorm:"column:phone;"`
	LastName  string                  `json:"last_name" gorm:"column:last_name;"`
	FirstName string                  `json:"first_name" gorm:"column:first_name;"`
	UserId    int                     `json:"-" gorm:"column:user_id;"`
	AvatarId  int                     `json:"avatar_id" gorm:"column:avatar_id;"`
	Image     *imageModel.SimpleImage `json:"image" gorm:"preload:false;foreignKey:AvatarId;"`
}

func (Profile) GetTableName() string { return ProfileRegister{}.GetTableName() }

func (p *Profile) Mark(isAdminOrOwner bool) {
	p.GetUID(common.OjbTypeProfile)
}

type ProfileUpdate struct {
	Email     string                  `json:"email" gorm:"column:email;"`
	Phone     string                  `json:"phone" gorm:"column:phone;"`
	LastName  string                  `json:"last_name" gorm:"column:last_name;"`
	FirstName string                  `json:"first_name" gorm:"column:first_name;"`
	AvatarId  int                     `json:"avatar_id" gorm:"column:avatar_id;"`
	Image     *imageModel.SimpleImage `json:"image" gorm:"preload:false;foreignKey:AvatarId;"`
}

func (ProfileUpdate) GetTableName() string { return ProfileUpdate{}.GetTableName() }
