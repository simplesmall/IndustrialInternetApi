package Response

import (
	"net/http"
)

type ResponseBody struct {
	Code int         `json:"code"`
	Msg  interface{}      `json:"msg"`
	Data interface{} `json:"data"`
}

//未找到响应体
func (res ResponseBody) NotFound() (result ResponseBody) {
	return ResponseBody{
		http.StatusNotFound,
		"Sorry,Not found",
		"",
	}
}

// 错误响应体
func (res ResponseBody) FailRes(err interface{}) (result ResponseBody) {
	return ResponseBody{
		http.StatusInternalServerError,
		"Sometthing wrong in the server end",
		err,
	}
}

// 错误响应体 + msg
func (res ResponseBody) FailResWithMsg(msg string, err interface{}) (result ResponseBody) {
	return ResponseBody{
		http.StatusInternalServerError,
		msg,
		err,
	}
}

// 错误响应体 + msg  + Code
func (res ResponseBody) FailResWithMsgCode(code int, msg string, err interface{}) (result ResponseBody) {
	return ResponseBody{
		code,
		msg,
		err,
	}
}

//带data 正常响应体
func (res ResponseBody) OKResult(data interface{}) (result ResponseBody) {
	result.Code = http.StatusOK
	result.Msg = "Normal......"
	result.Data = data
	return result
}

//带data 正常响应体 + msg
func (res ResponseBody) OKResultWithMsg(msg interface{}, data interface{}) (result ResponseBody) {
	result.Code = http.StatusOK
	result.Msg = msg
	result.Data = data
	return result
}
