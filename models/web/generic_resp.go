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
