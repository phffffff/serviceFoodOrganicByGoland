package common

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
)

type AppError struct {
	StatusCode int    `json:"status_code"`
	RootErr    error  `json:"-"`
	Message    string `json:"message"`
	Log        string `json:"log"`
	Key        string `json:"custom_code"`
}

func NewErrorResponse(
	rootErr error,
	msg string,
	log string,
	key string,
) *AppError {
	return &AppError{
		StatusCode: http.StatusBadRequest,
		RootErr:    rootErr,
		Message:    msg,
		Log:        log,
		Key:        key,
	}
}

func NewFullErrorResponse(
	statusCode int,
	rootErr error,
	msg string,
	log string,
	key string,
) *AppError {
	return &AppError{
		StatusCode: statusCode,
		RootErr:    rootErr,
		Message:    msg,
		Log:        log,
		Key:        key,
	}
}

func NewCustomError(root error, msg string, key string) *AppError {
	if root != nil {
		return NewErrorResponse(root, msg, root.Error(), key)
	}
	return NewErrorResponse(errors.New(msg), msg, msg, key)
}

func (e *AppError) RootError() error {
	if err, isOk := e.RootErr.(*AppError); isOk {
		return err.RootError()
	}
	return e.RootErr
}

func (e *AppError) Error() string {
	return e.RootErr.Error()
}

func ErrDB(err error) *AppError {
	return NewFullErrorResponse(
		http.StatusInternalServerError,
		err,
		MsgErrDb,
		err.Error(),
		ErrDBKey)
}

// Lỗi logic or syntax
func ErrInternal(err error) *AppError {
	return NewFullErrorResponse(
		http.StatusInternalServerError,
		err,
		MsgErrSv,
		err.Error(),
		ErrInternalKey)
}

func ErrInvalidRequest(err error) *AppError {
	return NewErrorResponse(err, MsgInvalidReq, err.Error(), ErrInvalidRequestKey)
}

func ErrCannotCRUDEntity(entity string, crud string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("cannot %s %s", strings.ToLower(crud), strings.ToLower(entity)),
		fmt.Sprintf(
			"ErrCannot%s%s",
			strings.ToTitle(strings.ToLower(crud)),
			strings.ToTitle(strings.ToLower(entity)),
		),
	)
}

func ErrRecordNotFound(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("%s not found", strings.ToLower(entity)),
		fmt.Sprintf("Err%sNotFound", strings.ToTitle(strings.ToLower(entity))),
	)
}

func ErrEntityDeleted(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("%s deleted", entity),
		fmt.Sprintf("Err%sDeleted", strings.ToTitle(strings.ToLower(entity))),
	)
}

func ErrorNoPermission(err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("you have no permission"),
		fmt.Sprintf("ErrNoPermission"),
	)
}
