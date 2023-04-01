package common

type AppResponse struct {
	Data   interface{}
	Filter interface{}
	Paging interface{}
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
