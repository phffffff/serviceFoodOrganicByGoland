package aboutModel

import (
	"go_service_food_organic/common"
)

const (
	EntityName = "About"
)

type About struct {
	common.SQLModel `json:",inline"`
	Title           string `json:"title" gorm:"column:title"`
	Description     string `json:"description" gorm:"column:description"`
	Logo            string `json:"logo" gorm:"column:logo"`
}

func (About) TableName() string { return "about_us" }

func (ab *About) Mask(isAdminOrOwner bool) {
	ab.GetUID(common.OjbTypeAbout)
}

type AboutUpdate struct {
	Title       string `json:"title" gorm:"column:title;"`
	Description string `json:"description" gorm:"column:description;"`
	Logo        string `json:"logo" gorm:"column:logo;"`
	Status      int    `json:"status" gorm:"column:status;"`
}

func (AboutUpdate) TableName() string { return About{}.TableName() }
