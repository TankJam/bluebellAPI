package controller

import (
	"bluebellAPI/dao/mysql"
	"bluebellAPI/logic"
	"bluebellAPI/models"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

/*
	user 请求数据接收与响应数据返回的处理
*/

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
	if err := logic.SignUp(p); err != nil {
		zap.L().Error("logic.SignUp failed", zap.Error(err))
		if errors.Is(err, mysql.ErrorUserExist){
			ResponseError(c, CodeUserExist)
			return
		}
		ResponseError(c, CodeServerBusy)
		return
	}

	// 3.返回响应
	ResponseSuccess(c, nil)
}

// LoginHandler 处理登录请求的函数
func LoginHandler(c *gin.Context){
	//1.获取请求的参数
	p := new(models.ParamLogin)

	// 校验json参数 c.ShouldBindJSON
	if err := c.ShouldBindJSON(p); err != nil{
		// 请求参数错误处理  zap.Error(err): 错误数据
		zap.L().Error("Login with invalid param error", zap.Error(err))
		// 判断参数是否是validator.ValidationErrors类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)  // 不是参数校验错误，则返回请求参数错误
			return
		}
		// 参数校验错误，则返回 格式化后 请求参数错误
		/// errs.Translate(trans)) 将英文的错误信息翻译成中文
		/// trans 是翻译器
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}

	// 2.业务逻辑处理
	token, err := logic.Login(p)
	if err != nil {
		//zap.L().Error()
		zap.L().Error("logic.Login failed", zap.String("username", p.Username), zap.Error(err))
		// 若用户不存在，则返回不存在的错误信息
		if errors.Is(err, mysql.ErrorUserNotExist) {
			ResponseError(c, CodeUserExist)
			return
		}
	}

	// 3.校验通过，返回登录成功以后的认证信息
	ResponseSuccess(c, token)
}