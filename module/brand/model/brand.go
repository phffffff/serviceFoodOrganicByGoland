package brandModel

import (
	"go_service_food_organic/common"
	foodModel "go_service_food_organic/module/food/model"
)

const (
	EntityName = "Brand"
)

type Brand struct {
	common.SQLModel `json:",inline"`
	Name            string            `json:"name" gorm:"column:name"`
	Description     string            `json:"description" gorm:"column:description"`
	Logo            string            `json:"logo" gorm:"column:logo"`
	Foods           []*foodModel.Food `json:"foods" gorm:"preload:false;foreignKey:BrandId;references:Id"`
}

func (Brand) TableName() string { return "brands" }

func (b *Brand) Mask(isAdminOrOwner bool) {
	b.GetUID(common.OjbTypeCategory)
}

type BrandCreate struct {
	common.SQLModel `json:",inline"`
	Name            string `json:"name" gorm:"column:name;"`
	Description     string `json:"description" gorm:"column:description;"`
	Logo            string `json:"logo" gorm:"column:logo;"`
}

func (BrandCreate) TableName() string { return Brand{}.TableName() }

func (b *BrandCreate) Mask(isAdminOrOwner bool) {
	b.GetUID(common.OjbTypeCategory)
}

type BrandUpdate struct {
	Name        string `json:"name" gorm:"column:name;"`
	Description string `json:"description" gorm:"column:description;"`
	Logo        string `json:"logo" gorm:"column:logo;"`
	Status      int    `json:"status" gorm:"column:status;"`
}

func (BrandUpdate) TableName() string { return Brand{}.TableName() }
