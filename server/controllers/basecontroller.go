package controllers

import (
	. "github.com/fishedee/encoding"
	. "github.com/fishedee/language"
	. "github.com/fishedee/web"
)

type BaseController struct {
	Controller
}

type BaseControllerData struct {
	Code int
	Data interface{}
	Msg  string
}

// 自动加载页面
func (this *BaseController) AutoRender(data interface{}, viewname string) {
	result := BaseControllerData{}
	resultError, isExcetion := data.(Exception)

	if isExcetion {
		//带错误码的error
		result.Code = resultError.GetCode()
		result.Msg = resultError.GetMessage()
		result.Data = nil
	} else {
		//正常返回
		result.Code = 0
		result.Data = data
		result.Msg = ""
	}

	resultString, err := EncodeJson(result)
	if err != nil {
		panic(err)
	}

	this.Ctx.Write(resultString)

	//exception, isException := data.(Exception)

	//var result struct {
	//	Code int         `json:"code"`
	//	Msg  string      `json:"msg"`
	//	Data interface{} `json:"data"`
	//}

	//if isException {
	//	result.Code = exception.GetCode()
	//	result.Msg = exception.GetMessage()
	//	result.Data = ""
	//} else {
	//	result.Code = 0
	//	result.Msg = ""
	//	result.Data = data
	//}

	//resultByte, err := json.Marshal(result)
	//if err != nil {
	//	panic(err)
	//}
	//this.Ctx.Write(resultByte)
}
