package utils

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

type HttpResult struct {
	Code int32    `json:"code"`
	Msg  string   `json:"msg"`
	Data struct{} `json:"data"`
}

func BuildOk(resp http.ResponseWriter) {
	BuildResponse("ok", resp)
}

func VerifyRequest(urlStr string) string {
	//查找字符串的位置
	questionIndex := strings.Index(urlStr, "?")
	//判断是否存在/符号
	cutIndex := strings.Index(urlStr, "/")
	//打散成数组
	rs := []rune(urlStr)
	//用于存储请求的地址切割
	requestSlice := make([]string, 0, 0)
	//用于存储请求的参数字典
	parameterDict := make(map[string]string)
	//请求地址
	requestStr := ""
	//参数地址
	parameterStr := ""
	//判断是否存在 ?
	if questionIndex != -1 {
		//判断url的长度
		parameterStr = string(rs[questionIndex+1 : len(urlStr)])
		requestStr = string(rs[0:questionIndex])
		//参数数组
		parameterArray := strings.Split(parameterStr, "&")
		//生成参数字典
		for i := 0; i < len(parameterArray); i++ {
			str := parameterArray[i]
			if len(str) > 0 {
				tem := strings.Split(str, "=")
				if len(tem) > 0 && len(tem) == 1 {
					parameterDict[tem[0]] = ""
				} else if len(tem) > 1 {
					parameterDict[tem[0]] = tem[1]
				}
			}
		}
	} else {
		requestStr = urlStr
	}

	//判断是否存在 /
	if cutIndex == -1 {
		requestSlice = append(requestSlice, requestStr)
	} else {
		//按 / 切割
		requestArray := strings.Split(requestStr, "/")
		for i := 0; i < len(requestArray); i++ {
			//判断第一个字符
			if i == 0 {
				//判断第一个字符串是否为空
				if len(requestArray[i]) != 0 {
					requestSlice = append(requestSlice, requestArray[i])
				}
			} else {
				requestSlice = append(requestSlice, requestArray[i])
			}
		}

	}

	if parameterDict["key"] == "yuki" {
		return requestStr
	} else {
		return ""
	}
}

func BuildResponse(message string, resp http.ResponseWriter) {
	var result HttpResult
	result.Msg = message
	result.Code = 0
	data, _ := json.Marshal(result)
	_, _ = io.WriteString(resp, string(data))
}
