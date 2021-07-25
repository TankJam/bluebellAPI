package controller

import (
	"bluebellAPI/models"
	"bluebellAPI/dao/mysql"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

// SignUpHandler 处理注册请求的函数
func SignUpHandler(c *gin.Context) {
	// 1.获取参数和参数校验
	p := new(models.ParamSignUp)
	if err := c.ShouldBindJSON(p); err != nil {
		// 请求参数错误则记录日志, 并直接返回响应
		zap.L().Error("SignUp with invalid param", zap.Error(err))
		// 判断err是不是validator.ValidationErrors类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			// 若不是参数校验错误，则返回请求参数错误
			ResponseError(c, CodeInvalidParam)
		}
		// 若是 校验参数错误，则 提示请求参数错误，并且 去除提示信息，返回自定义信息
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}

	// 2.业务处理
	// TODO: 2021-07-26 start
	if err := logic.SignUp(p); err != nil {
		zap.L().Error("logic.SignUp failed", zap.Error(err))
		if errors.Is(err, mysql.ErrorUserExist){
			ResponseError(c, CodeUserExist)
			return
		}
		ResponseError(c, CodeServerBusy)
	}
}
