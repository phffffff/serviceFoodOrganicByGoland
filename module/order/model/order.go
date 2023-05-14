package orderModel

import (
	"go_service_food_organic/common"
	orderDetailModel "go_service_food_organic/module/order_detail/model"
	userModel "go_service_food_organic/module/user/model"
)

const (
	EntityName = "Order"
)

type Order struct {
	common.SQLModel `json:",inline"`
	UserId          int                             `json:"-" gorm:"column:user_id"`
	UserFakeId      *common.UID                     `json:"user_id" gorm:"-"`
	Users           *userModel.User                 `json:"users" gorm:"preload:false;foreignKey:Id;references:UserId"`
	TotalPrice      float32                         `json:"total_price" gorm:"column:total_price"`
	State           string                          `json:"state" gorm:"column:state"`
	OrderDetails    []*orderDetailModel.OrderDetail `json:"order_details" gorm:"preload:false;foreignKey:OrderId;references:Id"`
}

func (Order) TableName() string { return "orders" }

func (o *Order) Mark(isAdminOrOwner bool) {
	o.GetUID(common.OjbTypeOrder)
	o.GetUserUID(common.OjbTypeOrder)
}

func (o *Order) GetUserUID(objType int) {
	uid := common.NewUID(uint32(o.UserId), objType, 1)
	o.UserFakeId = &uid
}

type OrderCreate struct {
	common.SQLModel `json:",inline"`
	UserId          int     `json:"-" gorm:"column:user_id"`
	UserFakeId      string  `json:"user_id" gorm:"-"`
	TotalPrice      float32 `json:"total_price" gorm:"column:total_price"`
}

func (OrderCreate) TableName() string { return Order{}.TableName() }

func (oc *OrderCreate) Mark(isAdminOrOwner bool) {
	oc.GetUID(common.OjbTypeOrder)
}

type OrderUpdate struct {
	TotalPrice float32 `json:"total_price" gorm:"column:total_price"`
	State      string  `json:"state" gorm:"column:state"`
	Status     int     `json:"status" gorm:"column:status"`
}

func (OrderUpdate) TableName() string { return Order{}.TableName() }
