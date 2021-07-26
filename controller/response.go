package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// ResponseData 响应信息结构体
type ResponseData struct {
	Code ResCode     `json:"code"`
	Msg  interface{} `json:"msg"`
	Data interface{} `json:"data"`
}

// ResponseError  返回错误响应函数
func ResponseError(c *gin.Context, code ResCode) { // code ResCode 错误码
	c.JSON(http.StatusOK, &ResponseData{
		Code: code,       // 返回对应的错误状态吗
		Msg:  code.Msg(), // 返回对应的错误信息
		Data: nil,        // 请求错误没有数据返回
	})
}


// ResponseErrorWithMsg 返回错误状态码，并返回自定义信息
func ResponseErrorWithMsg(c *gin.Context, code ResCode, msg interface{}) {
	c.JSON(http.StatusOK, &ResponseData{
		Code: code,
		Msg:  msg,
		Data: nil,
	})
}


// ResponseSuccess 请求成功返回响应
func ResponseSuccess(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, &ResponseData{
		Code: CodeSuccess,
		Msg:  CodeSuccess.Msg(),
		Data: data,
	})
}