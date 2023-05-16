package provinceModel

import "go_service_food_organic/common"

const (
	EntityName = "Province"
)

type Province struct {
	common.SQLModel `json:",inline"`
	Title           string `json:"title" gorm:"column:title"`
}

func (Province) TableName() string { return "provinces" }

func (p *Province) Mask(isAdminOrOwner bool) {
	p.GetUID(common.OjbTypeProvinces)
}
