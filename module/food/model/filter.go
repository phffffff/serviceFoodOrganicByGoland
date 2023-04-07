package foodModel

type Filter struct {
	Price   float32 `json:"price" form:"price"'`
	BrandId int     `json:"brand_id" form:"brand_id"`
	Status  int     `json:"status" form:"status"`
}
