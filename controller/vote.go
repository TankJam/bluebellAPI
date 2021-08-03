package controller

import (
	"bluebellAPI/logic"
	"bluebellAPI/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

// PostVoteHandler 投票分发函数
func PostVoteHandler(c *gin.Context) {
	// 参数校验
	p := new(models.ParamVoteData)

	fmt.Println(c.Params)

	if err := c.ShouldBindJSON(p); err != nil {
		// 错误类型断言，判断是否是参数校验报错
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			// 返回正常错误
			ResponseError(c, CodeInvalidParam)
			return
		}
		// 翻译并去除掉错误提示中的结构体标识
		errData := removeTopStruct(errs.Translate(trans))
		ResponseErrorWithMsg(c, CodeInvalidParam, errData)
		return
	}

	// 获取当前用户请求的用户id
	userID, err := getCurrentUser(c)
	if err != nil {
		ResponseError(c, CodeNeedLogin)
		return
	}

	// 具体投票的业务逻辑
	if err := logic.VoteForPost(userID, p); err != nil {
		zap.L().Error("logic.VoteForPost() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	ResponseSuccess(c, nil)
}
