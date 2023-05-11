package orderModel

import "go_service_food_organic/common"

const (
	EntityName = "Order"
)

type Order struct {
	common.SQLModel `json:",inline"`
	UserId          int         `json:"-" gorm:"column:user_id"`
	UserFakeId      *common.UID `json:"user_id" gorm:"-"`
	TotalPrice      float32     `json:"total_price" gorm:"column:total_price"`
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
