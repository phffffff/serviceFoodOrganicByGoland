package userModel

import (
	"go_service_food_organic/common"
)

type User struct {
	common.SQLModel `json:",inline"`
	Email           string `json:"email" gorm:"column:email;"`
	FbId            string `json:"fb_id,omitempty" gorm:"column:fb_id;"`
	GgId            string `json:"gg_id,omitempty" gorm:"column:gg_id;"`
	Password        string `json:"-" gorm:"column:password;"`
	Salt            string `json:"-" gorm:"column:salt;"`
	Phone           string `json:"Phone" gorm:"column:phone;"`
	Role            string `json:"role" gorm:"column:role;default:'user'"`
}

func (User) TableName() string { return "users" }

func (u *User) Mark(isAdminOrOwner bool) {
	u.GetUID(common.OjbTypeUser)
}

func (u *User) GetUserId() int   { return u.Id }
func (u *User) GetEmail() string { return u.Email }
func (u *User) GetRole() string  { return u.Role }

type UserPasswordUpdate struct {
	common.SQLModel `json:",inline"`
	Password        string `json:"password" gorm:"-"`
	NewPassword     string `json:"new_password" gorm:"column:password;"`
	ReNewPassword   string `json:"re_new_password" gorm:"-"`
}

func (UserPasswordUpdate) TableName() string { return User{}.TableName() }

func (u *UserPasswordUpdate) Mark(isAdminOrOwner bool) {
	u.GetUID(common.OjbTypeUser)
}

type UserRegister struct {
	common.SQLModel `json:",inline"`
	Email           string `json:"email" gorm:"column:email;"`
	Password        string `json:"password" gorm:"column:password;"`
	RePassword      string `json:"re_password" gorm:"-;"`
	Salt            string `json:"salt" gorm:"column:salt;"`
	Phone           string `json:"phone" gorm:"column:phone;"`
	FbId            string `json:"fb_id,omitempty" gorm:"column:fb_id;"`
	GgId            string `json:"gg_id,omitempty" gorm:"column:gg_id;"`
}

func (u *UserRegister) Mark(isAdminOrOwner bool) {
	u.GetUID(common.OjbTypeUser)
}

func (UserRegister) TableName() string { return User{}.TableName() }

type UserLogin struct {
	Email    string `json:"email" form:"email;"`
	Password string `json:"password" form:"password;"`
}

func (UserLogin) TableName() string { return User{}.TableName() }

func ErrorEmailOrPasswordInvalid(err error) *common.AppError {
	return common.NewCustomError(err, MsgEmailOrPasswordInvalid, ErrEmailOrPasswordInvalid)
}
func ErrorRePassInvalid(err error) *common.AppError {
	return common.NewCustomError(err, ErrRePassInvalid, MsgRePassInvalid)
}
func ErrorNewPassInvalid(err error) *common.AppError {
	return common.NewCustomError(err, ErrNewPassInvalid, MsgNewPassInvalid)
}

const (
	EntityName = "User"

	ErrEmailOrPasswordInvalid = "ErrEmailOrPasswordInvalid"
	MsgEmailOrPasswordInvalid = "email or password invalid"

	ErrRePassInvalid = "ErrRePassInvalid"
	MsgRePassInvalid = "re-pass invalid"

	ErrNewPassInvalid = "ErrNewPassInvalid"
	MsgNewPassInvalid = "new password invalid"
)
