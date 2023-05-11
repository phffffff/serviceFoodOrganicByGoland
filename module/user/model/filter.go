package userModel

type Filter struct {
	Status []int `json:"status" form:"status"`
}
