package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"strconv"
)

/*
	处理请求的控制器
*/

const CtxUserIDKey = "userID"

var ErrorUserNotLogin = errors.New("用户未登录")

// getCurrentUser 获取当前登录的用户ID
func getCurrentUser(c *gin.Context) (userID int64, err error) {
	// 只要请求通过jwt，就可以从c上下文中获取userID
	uid, ok := c.Get(CtxUserIDKey)
	if !ok {
		err = ErrorUserNotLogin // 用户未登录
		return
	}

	userID, ok = uid.(int64) // 判断uid是否是int64
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	return
}

// getPageInfo 获取分页参数
func getPageInfo(c *gin.Context) (int64, int64) {
	// 页数、条数
	pageStr := c.Query("page")
	sizeStr := c.Query("size")

	var (
		page int64
		size int64
		err  error
	)

	page, err = strconv.ParseInt(pageStr, 10, 64)
	// 没有 page 默认第 1 页
	if err != nil {
		page = 1
	}

	// 没有 size 默认 10
	size, err = strconv.ParseInt(sizeStr, 10, 64)
	if err != nil {
		size = 10
	}
	return page, size
}
