package orderDetailModel

import "go_service_food_organic/common"

const (
	EntityName = "Order Detail"
)

type OrderDetail struct {
	common.SQLModel `json:",inline"`
	OrderId         int         `json:"-" gorm:"column:order_id"`
	OrderFakeId     *common.UID `json:"order_id" gorm:"-"`
	FoodId          int         `json:"-" gorm:"column:food_id"`
	FoodFakeId      *common.UID `json:"food_id" gorm:"-"`
	Quantity        int         `json:"quantity" gorm:"column:quantity"`
	Price           float32     `json:"price" gorm:"column:price"`
}

func (OrderDetail) TableName() string {
	return "order_details"
}

func (od *OrderDetail) Mark(isAdminOrOwner bool) {
	od.GetUID(common.OjbTypeOrderDetail)
	od.GetOrderUID(common.OjbTypeOrder)
	od.GetFoodUID(common.OjbTypeFood)
}
func (od *OrderDetail) GetOrderUID(objType int) {
	uid := common.NewUID(uint32(od.OrderId), objType, 1)
	od.OrderFakeId = &uid
}
func (od *OrderDetail) GetFoodUID(objType int) {
	uid := common.NewUID(uint32(od.OrderId), objType, 1)
	od.FoodFakeId = &uid
}

type OrderDetailCreate struct {
	common.SQLModel `json:",inline"`
	OrderId         int     `json:"-" gorm:"column:order_id"`
	OrderFakeId     string  `json:"order_id" gorm:"-"`
	FoodId          int     `json:"-" gorm:"column:food_id"`
	FoodFakeId      string  `json:"food_id" gorm:"-"`
	Quantity        int     `json:"quantity" gorm:"column:quantity"`
	Price           float32 `json:"price" gorm:"column:price"`
}
