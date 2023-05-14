package cartModel

import (
	"go_service_food_organic/common"
	"time"
)

const (
	EntityName = "Cart"
)

type Cart struct {
	UserId     int        `json:"-" gorm:"column:user_id"`
	FoodId     int        `json:"-" gorm:"column:food_id"`
	FoodFakeId string     `json:"food_id" gorm:"-"`
	Quantity   int        `json:"quantity" gorm:"column:quantity"`
	Price      float32    `json:"price" gorm:"column:price"`
	CreatedAt  *time.Time `json:"created_at" gorm:"column:created_at;"`
	UpdatedAt  *time.Time `json:"updated_at" gorm:"column:updated_at;"`
}

func (Cart) TableName() string { return "carts" }

type CartLst struct {
	UserId     int         `json:"-" gorm:"column:user_id"`
	FoodId     int         `json:"-" gorm:"column:food_id"`
	FoodFakeId *common.UID `json:"food_id" gorm:"-"`
	Quantity   int         `json:"quantity" gorm:"column:quantity"`
	Price      float32     `json:"price" gorm:"column:price"`
	CreatedAt  *time.Time  `json:"created_at" gorm:"column:created_at;"`
	UpdatedAt  *time.Time  `json:"updated_at" gorm:"column:updated_at;"`
}

func (cart *CartLst) GetFoodUID(OjbType int) {
	uid := common.NewUID(uint32(cart.FoodId), OjbType, 1)
	cart.FoodFakeId = &uid
	//log.Print(uid.String())
}

func (cart *CartLst) Mask(isAdminOrOwner bool) {
	cart.GetFoodUID(common.OjbTypeFood)
}
