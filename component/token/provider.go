package token

import (
	"errors"
	"go_service_food_organic/common"
	"time"
)

type Provider interface {
	Generate(data TokenPayload, expiry int) (*Token, error)
	Validate(token string) (*TokenPayload, error)
}

var (
	ErrorNotFound = common.NewCustomError(
		errors.New("token not found"),
		"token not found",
		"ErrNoteFound",
	)
	ErrorEncodingToken = common.NewCustomError(
		errors.New("error encoding the token"),
		"error encoding the token",
		"ErrEncodingToken",
	)
	ErrorInvalidToken = common.NewCustomError(
		errors.New("error invalid token"),
		"error invalid token",
		"ErrInvalidToken",
	)
)

type Token struct {
	Token   string    `json:"token"`
	Created time.Time `json:"created"`
	Expiry  int       `json:"expiry"`
}

type TokenPayload struct {
	UserId int         `json:"-"`
	FakeId *common.UID `json:"user_id"`
	Role   string      `json:"role"`
}
