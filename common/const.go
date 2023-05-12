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

	MsgCannotCommit = "can not commit"
	ErrCannotCommit = "ErrCannotCommit"

	MsgCannotRollback = "can not roll back"
	ErrCannotRollback = "ErrCannotRollback"

	OjbTypeUser        = 1
	OjbTypeFood        = 2
	OjbTypeProfile     = 3
	OjbTypeImage       = 4
	OjbTypeOrder       = 5
	OjbTypeOrderDetail = 6

	OjbTypeCart = 7
)

type Requester interface {
	GetUserId() int
	GetEmail() string
	GetRole() string
}
