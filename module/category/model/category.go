package categoryModel

import (
	"go_service_food_organic/common"
	foodModel "go_service_food_organic/module/food/model"
)

const (
	EntityName = "Category"
)

type Category struct {
	common.SQLModel `json:",inline"`
	Name            string            `json:"name" gorm:"column:name"`
	Description     string            `json:"description" gorm:"description"`
	Icon            string            `json:"icon" gorm:"column:icon"`
	Foods           []*foodModel.Food `json:"foods" gorm:"many2many:info_food_categories;foreignKey:Id;joinForeignKey:CategoryId;References:Id;joinReferences:FoodId"`
}

func (Category) TableName() string { return "categories" }

func (c *Category) Mask(isAdminOrOwner bool) {
	c.GetUID(common.OjbTypeCategory)
}

type CategoryCreate struct {
	common.SQLModel `json:",inline"`
	Name            string `json:"name" gorm:"column:name;"`
	Description     string `json:"description" gorm:"column:description;"`
	Icon            string `json:"icon" gorm:"column:icon;"`
}

func (CategoryCreate) TableName() string { return Category{}.TableName() }

func (c *CategoryCreate) Mask(isAdminOrOwner bool) {
	c.GetUID(common.OjbTypeCategory)
}

type CategoryUpdate struct {
	Name        string `json:"name" gorm:"column:name;"`
	Description string `json:"description" gorm:"column:description;"`
	Icon        string `json:"icon" gorm:"column:icon;"`
	Status      int    `json:"status" gorm:"column:status;"`
}

func (CategoryUpdate) TableName() string { return Category{}.TableName() }
