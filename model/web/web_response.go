package web

type WebResponse struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

type WebResponseError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
