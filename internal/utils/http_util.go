package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

type HttpResult struct {
	Code int32    `json:"code"`
	Msg  string   `json:"msg"`
	Data struct{} `json:"data"`
}

func BuildOk(resp http.ResponseWriter) {
	BuildResponse("ok", resp)
}

func BuildResponse(message string, resp http.ResponseWriter) {
	var result HttpResult
	result.Msg = message
	result.Code = 0
	data, _ := json.Marshal(result)
	_, _ = io.WriteString(resp, string(data))
}
