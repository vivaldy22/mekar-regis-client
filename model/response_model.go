package model

type JSONResponse struct {
	Success bool `json:"success"`
	Message string `json:"message"`
	Data interface{} `json:"data"`
}

type JSONFailResponse struct {
	Success bool `json:"success"`
	Message string `json:"message"`
	ErrorCode int `json:"error_code"`
	Data interface{} `json:"data"`
}