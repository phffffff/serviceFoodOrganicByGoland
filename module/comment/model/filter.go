package commentModel

type Filter struct {
	Status    []int `json:"status" form:"status"`
	ProfileId int   `json:"profile_id" form:"profile_id"`
	NewId     int   `json:"new_id" form:"new_id"`
}
