package common

const (
	Admin = "admin"

	CurrentUser = "user"

	Get    = "GET"
	List   = "LIST"
	Create = "CREATE"
	Update = "UPDATE"
	Delete = "DELETE"

	MsgErrDb = "something went wrong with DB"
	ErrDBKey = "DB_ERROR"

	MsgErrSv       = "something went wrong with Server"
	ErrInternalKey = "ErrInternal"

	MsgInvalidReq        = "invalid request"
	ErrInvalidRequestKey = "ErrInvalidRequest"

	OjbTypeUser    = 1
	OjbTypeFood    = 2
	OjbTypeProfile = 3
)

type Requester interface {
	GetUserId() int
	GetEmail() string
	GetRole() string
}
