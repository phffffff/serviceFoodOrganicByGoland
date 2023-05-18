package addressModel

import (
	"go_service_food_organic/common"
	provinceModel "go_service_food_organic/module/province/model"
)

const (
	EntityName = "Address"
)

type Address struct {
	common.SQLModel `json:",inline"`
	ProfileId       int                     `json:"-" gorm:"column:profile_id"`
	ProfileFakeId   *common.UID             `json:"profile_id" gorm:"-"`
	ProvinceId      int                     `json:"-" gorm:"column:province_id"`
	ProvinceFakeId  *common.UID             `json:"province_id" gorm:"-"`
	Provinces       *provinceModel.Province `json:"provinces" gorm:"preload:false;foreignKey:Id;references:ProvinceId"`
	Title           string                  `json:"title" gorm:"column:title"`
	Addr            string                  `json:"addr" gorm:"column:addr"`
	IsDefault       string                  `json:"is_default" gorm:"column:is_default"`
	ZipId           int                     `json:"zip_id" gorm:"column:zip_id"`
}

func (Address) TableName() string { return "user_addresses" }

func (a *Address) Mask(isAdminOrOwner bool) {
	a.GetUID(common.OjbTypeAddress)
	a.GetProfileUID(common.OjbTypeProfile)
	a.GetProvinceUID(common.OjbTypeProvinces)
}
func (a *Address) GetProvinceUID(objType int) {
	uid := common.NewUID(uint32(a.ProvinceId), objType, 1)
	a.ProfileFakeId = &uid
}
func (a *Address) GetProfileUID(objType int) {
	uid := common.NewUID(uint32(a.ProfileId), objType, 1)
	a.ProfileFakeId = &uid
}

type AddressCreate struct {
	common.SQLModel `json:",inline"`
	ProfileId       int    `json:"-" gorm:"column:profile_id"`
	ProfileFakeId   string `json:"profile_id" gorm:"-"`
	ProvinceId      int    `json:"-" gorm:"column:province_id"`
	ProvinceFakeId  string `json:"province_id" gorm:"-"`
	Title           string `json:"title" gorm:"column:title"`
	Addr            string `json:"addr" gorm:"column:addr"`
	IsDefault       string `json:"is_default" gorm:"column:is_default"`
	ZipId           int    `json:"zip_id" gorm:"column:zip_id"`
}

func (AddressCreate) TableName() string { return Address{}.TableName() }

func (ac *AddressCreate) Mask(isAdminOrOwner bool) {
	ac.GetUID(common.OjbTypeAddress)
}

type AddressUpdate struct {
	ProfileId      int    `json:"-" gorm:"column:profile_id"`
	ProfileFakeId  string `json:"profile_id" gorm:"-"`
	ProvinceId     int    `json:"-" gorm:"column:province_id"`
	ProvinceFakeId string `json:"province_id" gorm:"-"`
	Title          string `json:"title" gorm:"column:title"`
	Addr           string `json:"addr" gorm:"column:addr"`
	IsDefault      string `json:"is_default" gorm:"column:is_default"`
	Status         int    `json:"status" gorm:"column:status"`
	ZipId          int    `json:"zip_id" gorm:"column:zip_id"`
}

func (AddressUpdate) TableName() string { return Address{}.TableName() }
