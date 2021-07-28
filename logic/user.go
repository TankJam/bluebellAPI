package logic

import (
	"bluebellAPI/dao/mysql"
	"bluebellAPI/models"
	"bluebellAPI/pkg/jwt"
	"bluebellAPI/pkg/snowflake"
	"fmt"
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
	fmt.Println(userID)

	// 3.构造一个User实例
	user := &models.User{
		UserID:   userID,
		Username: p.Username,
		Password: p.Password,
	}

	// 4.保存数据到数据库
	return mysql.InsertUser(user)

}

// Login 登录的业务逻辑处理
func Login(p *models.ParamLogin) (token string, err error) {
	user := &models.User{
		Username: p.Username,
		Password: p.Password,
	}
	// 1.传入指针，就能拿到user.UserID
	/// 若登录成功则返回报错信息
	if err := mysql.Login(user); err != nil {
		return "", err
	}

	/// 登录成功则返回token
	return jwt.GenToken(user.UserID, user.Username)
}
