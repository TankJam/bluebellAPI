package mysql

import "errors"

/*
	mysql 错误状态吗
*/

var (
	ErrorUserExist       = errors.New("用户已存在")
	ErrorUserNotExist    = errors.New("用户不存在")
	ErrorInvalidPassword = errors.New("用户名或密码错误")
	//ErrorInvalidID       = errors.New("无效ID")
)
