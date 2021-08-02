package controller

import (
	"bluebellAPI/logic"
	"bluebellAPI/models"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
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

// GetPostListHandler 查询所有帖子 V1
func GetPostListHandler(c *gin.Context) {
	// 1.获取分页参数
	page, size := getPageInfo(c)

	// 2.获取数据
	data, err := logic.GetPostList(page, size)
	if err != nil {
		zap.L().Error("logic.GetPostList() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	// 3.返回响应
	ResponseSuccess(c, data)
}

// GetPostListHandler2 根据时间查询所有帖子 V2
func GetPostListHandler2(c *gin.Context){
	/*
	- 根据前端传来的参数动态的获取帖子列表
	- 按创建时间排序 或者 按照 分数排序
		1.获取请求的 query string 参数
		2.去redis查询id列表
		3.根据id去数据库查询帖子详细信息
	*/
	// 初始化结构体时指定初始参数
	p := &models.ParamPostList{
		Page:        1,
		Size:        10,
		Order:       models.OrderTime,  // magic string
	}

	// 校验参数
	if err := c.ShouldBindQuery(p); err != nil {
		zap.L().Error("GetPostListHandler2 with invalid params", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	// TODO: 2021-08-03
}

// GetPostDetailHandler 根据id获取post详情数据
func GetPostDetailHandler(c *gin.Context){
	// 1.获取id参数
	pidStr := c.Param("id")
	pid, err := strconv.ParseInt(pidStr, 10, 64)
	if err != nil {
		zap.L().Error("get post detail with invalid param", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}

	// 2.根据id去除帖子数据
	data, err := logic.GetPostById(pid)
	if err != nil {
		zap.L().Error("logic.GetPostById(pid)", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	// 3.返回响应
	ResponseSuccess(c, data)
}