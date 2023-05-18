package addressModel

type Filter struct {
	Status     []int  `json:"status" form:"status"`
	ProfileId  int    `json:"profile_id" form:"profile_id"`
	ProvinceId int    `json:"province_id" form:"province_id"`
	IsDefault  string `json:"is_default" form:"is_default"`
}
