package model

type HttpResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

type GeneralListQuery struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}
