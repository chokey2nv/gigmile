package dtos

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
type Response struct {
	Data any `json:"data"`
}
type PageOption struct {
	Skip  int64 `json:"skip"`
	Limit int64 `json:"limit"`
}
