package web

type PlainResp struct {
	Status string `json:"status"`
	Msg    string `json:"message"`
}

type WithDataResp struct {
	Status string `json:"status"`
	Msg    string `json:"message"`
	Data   any    `json:"data"`
}

func ErrWithMsg(msg string) map[string]string {
	return map[string]string{
		"status":  "errror",
		"message": msg,
	}
}
