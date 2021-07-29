package controller

import (
	"bluebellAPI/logic"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
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


// CommunityDetailHandler 根据id返回
func CommunityDetailHandler(c *gin.Context) {
	// 1.获取社区id
	idStr := c.Param("id")  // 获取url中的id对应的value
	// string转为int 10进制 64位
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ResponseError(c, CodeInvalidParam)  // 参数错误
		return
	}

	// 2.根据id获取社区详情数据
	data, err := logic.GetCommunityDetail(id)
	if err != nil {
		zap.L().Error("logic.GetCommunityList() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)  // 服务繁忙
		return
	}

	ResponseSuccess(c, data)
}

