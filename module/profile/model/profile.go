package profileModel

import "go_service_food_organic/common"

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
	Email     string `json:"email" gorm:"column:email;"`
	Phone     string `json:"salt" gorm:"column:phone;"`
	LastName  string `json:"last_name" gorm:"column:last_name;"`
	FirstName string `json:"first_name" gorm:"column:first_name;"`
	//Avatar
}

func (Profile) GetTableName() string { return ProfileRegister{}.GetTableName() }

func (p *Profile) Mark(isAdminOrOwner bool) {
	p.GetUID(common.OjbTypeProfile)
}
