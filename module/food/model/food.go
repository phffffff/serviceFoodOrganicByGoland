package foodModel

import (
	"go_service_food_organic/common"
	"go_service_food_organic/module/image_food/model"
)

const EntityName = "Food"

// modelList
type Food struct {
	common.SQLModel `json:",inline"`
	Name            string                      `json:"name" gorm:"column:name;"`
	Description     string                      `json:"description" gorm:"description;"`
	Price           float32                     `json:"price" gorm:"column:price;"`
	Count           int64                       `json:"count" gorm:"column:count;"`
	BrandId         int                         `json:"-" gorm:"column:brand_id;"`
	BrandFakeId     *common.UID                 `json:"brand_id" gorm:"-"`
	FoodImages      []*imageFoodModel.ImageFood `json:"food_images" gorm:"foreignKey:FoodId;preload:false;"`
}

func (Food) TableName() string { return "foods" }

func (f *Food) GetBrandUID(OjbType int) {
	uid := common.NewUID(uint32(f.BrandId), OjbType, 1)
	f.BrandFakeId = &uid
}
func (f *Food) Mark(isAdminOrOwner bool) {
	f.GetUID(common.OjbTypeFood)
	f.GetBrandUID(common.OjbTypeFood)
}

type FoodCreate struct {
	common.SQLModel `json:",inline"`
	Name            string  `json:"name" gorm:"column:name;" `
	Description     string  `json:"description" gorm:"description;"`
	Price           float32 `json:"price" gorm:"column:price;"`
	Count           int64   `json:"count" gorm:"column:count;"`
	BrandId         int     `json:"brand_id" gorm:"column:brand_id;"`
}

func (f *FoodCreate) Mark(isAdminOrOwner bool) {
	f.GetUID(common.OjbTypeFood)
}

func (data *FoodCreate) Validate() {
}

func (FoodCreate) TableName() string { return Food{}.TableName() }

type FoodUpdate struct {
	Name        string  `json:"name" gorm:"column:name;"`
	Description string  `json:"description" gorm:"description;"`
	Price       float32 `json:"price" gorm:"column:price;"`
	Count       int64   `json:"count" gorm:"column:count;"`
	BrandId     int     `json:"-" gorm:"column:brand_id;"`
	BrandFakeId string  `json:"brand_id" gorm:"-"`
}

func (FoodUpdate) TableName() string { return Food{}.TableName() }
