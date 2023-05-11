package orderDetailModel

type Filter struct {
	Status []int `json:"status" form:"status"`
}
