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
	FoodFakeId *string    `json:"food_id" gorm:"-"`
	Quantity   int        `json:"quantity" gorm:"column:quantity"`
	Price      float32    `json:"price" gorm:"column:price"`
	CreateAt   *time.Time `json:"create_at" gorm:"column:create_at"`
	UpdateAt   *time.Time `json:"update_at" gorm:"column:update_at"`
}

func (cart *Cart) GetUID() {
	foodFakeId := common.NewUID(uint32(cart.FoodId), common.OjbTypeFood, 1).String()
	cart.FoodFakeId = &foodFakeId
}

func (cart *Cart) Mask(isAdminOrOwner bool) {
	cart.GetUID()
}
func (Cart) TableName() string { return "carts" }
