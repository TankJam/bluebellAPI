package controller

import (
	"bluebellAPI/logic"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

/*
	----- 社区相关的Handler -----
*/

// CommunityHandler 请求响应处理
func CommunityHandler(c *gin.Context) {
	// 查询到所有的社区（community_id, community_name） 以列表的形式返回
	data, err := logic.GetCommunityList()
	if err != nil {
		zap.L().Error("logic.GetCommunityList() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)  // 不轻易把服务端报错暴露给外面，所以返回服务器繁忙
		return
	}
	ResponseSuccess(c, data)
}
