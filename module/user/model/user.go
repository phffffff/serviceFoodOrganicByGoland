package userModel

import (
	"go_service_food_organic/common"
)

type User struct {
	common.SQLModel `json:",inline"`
	Email           string `json:"email" gorm:"column:email;"`
	FbId            string `json:"fb_id" gorm:"column:fb_id;"`
	GgId            string `json:"gg_id" gorm:"column:gg_id;"`
	Password        string `json:"-" gorm:"column:password;"`
	Salt            string `json:"-" gorm:"column:salt;"`
	Phone           string `json:"salt" gorm:"column:phone;"`
	Role            string `json:"role" gorm:"column:role;default:'user'"`
}

func (User) GetTableName() string { return "users" }

func (u *User) Mark(isAdminOrOwner bool) {
	u.GetUID(common.OjbTypeUser)
}

func (u *User) GetUserId() int   { return u.Id }
func (u *User) GetEmail() string { return u.Email }
func (u *User) GetRole() string  { return u.Role }

type UserRegister struct {
	common.SQLModel `json:",inline"`
	Email           string `json:"email" gorm:"column:email;"`
	Password        string `json:"password" gorm:"column:password;"`
	Salt            string `json:"salt" gorm:"column:salt;"`
	Phone           string `json:"phone" gorm:"column:phone;"`
	FbId            string `json:"fb_id" gorm:"column:fb_id;"`
	GgId            string `json:"gg_id" gorm:"column:gg_id;"`
}

func (u *UserRegister) Mark(isAdminOrOwner bool) {
	u.GetUID(common.OjbTypeUser)
}

func (UserRegister) GetTableName() string { return User{}.GetTableName() }

type UserLogin struct {
	Email    string `json:"email" form:"email;"`
	Password string `json:"password" form:"password;"`
}

func (UserLogin) GetTableName() string { return User{}.GetTableName() }

func ErrorEmailOrPasswordInvalid(err error) *common.AppError {
	return common.NewCustomError(err, MsgEmailOrPasswordInvalid, ErrEmailOrPasswordInvalid)
}

const (
	EntityName = "User"

	ErrEmailOrPasswordInvalid = "ErrEmailOrPasswordInvalid"
	MsgEmailOrPasswordInvalid = "email or password invalid"
)
