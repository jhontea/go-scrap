package responses

import "net/http"

type UrlHttpResponse struct {
	Status     string      `json:"status"`
	StatusCode int         `json:"status_code"`
	Body       string      `json:"body"`
	Header     http.Header `json:"body"`
}

type Response struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
	Status  string      `json:"status"`
}
