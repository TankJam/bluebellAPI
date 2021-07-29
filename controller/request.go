package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
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
		err = ErrorUserNotLogin  // 用户未登录
		return
	}

	userID, ok = uid.(int64)  // 判断uid是否是int64
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	return
}