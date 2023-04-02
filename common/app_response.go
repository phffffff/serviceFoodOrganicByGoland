package common

type AppResponse struct {
	Data   interface{} `json:"data"`
	Filter interface{} `json:"filter,omitempty"`
	Paging interface{} `json:"paging,omitempty"`
}

func FullSuccessResponse(data, filter, paging interface{}) *AppResponse {
	return &AppResponse{
		Data:   data,
		Filter: filter,
		Paging: paging,
	}
}

func SimpleSuccessResponse(data interface{}) *AppResponse {
	return FullSuccessResponse(data, nil, nil)
}
