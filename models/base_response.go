package models

type BaseResponseError struct {
	ErrorCode int    `json:"error_code"`
	Message   string `json:"message"`
	Status    bool   `json:"status"`
}

type BaseResponseSuccess struct {
	Message string `json:"message"`
	Status  bool   `json:"status"`
	Data    any    `json:"data"`
}
