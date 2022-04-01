package utils

type (
	NormalResp struct {
		Status  string `json:"success"`
		Message string `json:"message"`
	}
	RespWithData struct {
		Status  string      `json:"status"`
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}
)
