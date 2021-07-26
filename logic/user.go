package logic

import (
	"bluebellAPI/dao/mysql"
	"bluebellAPI/models"
	"bluebellAPI/pkg/snowflake"
)

/*
	user 模块 业务逻辑的处理
*/

// SignUp 注册业务逻辑处理
func SignUp(p *models.ParamSignUp) (err error) {
	// 1、判断用户是否存在
	if err := mysql.CheckUserExist(p.Username); err != nil {
		return err
	}

	// 2.生成UID
	userID := snowflake.GenID()

	// 3.构造一个User实例
	user := &models.User{
		UserID:   userID,
		Username: p.Username,
		Password: p.Password,
	}

	// 4.保存数据到数据库
	return mysql.InsertUser(user)

}
