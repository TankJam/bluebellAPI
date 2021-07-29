package controller

import (
	"bluebellAPI/logic"
	"bluebellAPI/models"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// CreatePostHandler 新增帖子
func CreatePostHandler(c *gin.Context){
	// 1.获取参数以及参数的校验
	p := new(models.Post)

	if err := c.ShouldBindJSON(p); err != nil{
		zap.L().Debug("c.ShouldBindJSON(p) err", zap.Any("err", err))
		zap.L().Error("create post with invalid param")
		ResponseError(c, CodeInvalidParam)
		return
	}

	// 从 c 渠道当前发送请求的用户 userID
	userID, err := getCurrentUser(c)
	if err != nil {
		ResponseError(c, CodeNeedLogin)
		return
	}

	p.AuthorID = userID

	// 2.创建帖子
	if err := logic.CreatePost(p); err != nil {
		zap.L().Error("logic.CreatePost(p) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	// 3.返回响应
	ResponseSuccess(c, nil)
}
