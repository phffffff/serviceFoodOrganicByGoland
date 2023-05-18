package newModel

type Filter struct {
	Status []int `json:"status" form:"status"`
	Author int   `json:"author" form:"author"`
}
