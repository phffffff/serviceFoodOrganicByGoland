package foodModel

import (
	"go_service_food_organic/common"
)

const EntityName = "Food"

// modelList
type Food struct {
	common.SQLModel `json:",inline"`
	Name            string  `json:"name" gorm:"column:name;"`
	Description     string  `json:"description" gorm:"description;"`
	Price           float32 `json:"price" gorm:"column:price;"`
	Count           int64   `json:"count" gorm:"column:count;"'`
	//Image           *common.Image  `json:"image" gorm:"column:image;"`
	//Images          *common.Images `json:"images" gorm:"column:images;"`
	BrandId int `json:"brand_id" gorm:"column:brand_id;"`
}

func (Food) GetTableName() string { return "foods" }

func (f *Food) Mark(isAdminOrOwner bool) {
	f.GetUID(common.OjbTypeFood)
}

type FoodCreate struct {
	common.SQLModel `json:",inline"`
	Name            string  `json:"name" gorm:"column:name;" `
	Description     string  `json:"description" gorm:"description;"`
	Price           float32 `json:"price" gorm:"column:price;"`
	Count           int64   `json:"count" gorm:"column:count;"'`
}

func (f *FoodCreate) Mark(isAdminOrOwner bool) {
	f.GetUID(common.OjbTypeFood)
}

func (data *FoodCreate) Validate() {
}

func (FoodCreate) GetTableName() string { return Food{}.GetTableName() }
