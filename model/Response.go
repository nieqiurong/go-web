package model

type Response struct {
	BaseResponse
	Data interface{} `json:"data"`
}
