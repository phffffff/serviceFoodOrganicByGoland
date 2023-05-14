package infoFoodCategoryModel

import "go_service_food_organic/common"

type InfoFoodCategory struct {
	common.SQLModel `json:",inline"`
	FoodId          int         `json:"-" gorm:"column:food_id"`
	FoodFakeId      *common.UID `json:"food_id" gorm:"-"`
	CategoryId      int         `json:"-" gorm:"column:category_id"`
	CategoryFakeId  *common.UID `json:"category_id" gorm:"-"`
}

func (InfoFoodCategory) TableName() string { return "info_food_categories" }

func (ifc *InfoFoodCategory) Mask(isAdminOrOwner bool) {
	ifc.GetUID(common.OjbTypeInfoFoodCategoy)
	ifc.GetFoodUID(common.OjbTypeFood)
	ifc.GetCategoryUID(common.OjbTypeCategory)
}

func (ifc *InfoFoodCategory) GetFoodUID(ObjType int) {
	uid := common.NewUID(uint32(ifc.FoodId), ObjType, 1)
	ifc.FoodFakeId = &uid
}

func (ifc *InfoFoodCategory) GetCategoryUID(ObjType int) {
	uid := common.NewUID(uint32(ifc.CategoryId), ObjType, 1)
	ifc.CategoryFakeId = &uid
}
