package mysql

import (
	"bluebellAPI/models"
	"crypto/md5"
	"encoding/hex"
)

/*
	user 模块 数据层处理
*/

// secret 密码的加密 盐
var secret = "tank jam is very handsome!"

// CheckUserExist 判断用户是否存在
func CheckUserExist(username string) (err error) {
	sqlStr := `select count(user_id) from user where username = ?`
	var count int64
	// orm查询，若报错则sql执行错误
	if err := db.Get(&count, sqlStr, username); err != nil {
		return err
	}

	// count > 0 代表用户已存在
	if count > 0{
		return ErrorUserExist  // 返回用户已存在状态码
	}
	return
}

// InsertUser 插入注册的新用户
func InsertUser(user *models.User) (err error) {
	// 对密码进行加密
	user.Password = encryptPassword(user.Password)
	// 执行sql语句
	sqlStr := `insert into user(user_id, username, password) values(?,?,?)`
	_, err = db.Exec(sqlStr, user.UserID, user.Username, user.Password)
	return
}

// encryptPassword 对密码进行加密
func encryptPassword(oldPassword string) string {
	h := md5.New()
	h.Write([]byte(secret))  // 加盐
	return hex.EncodeToString(h.Sum([]byte(oldPassword)))
}

// Login 登录数据处理函数
func Login(user *models.User) (err error) {
	oldPassword := user.Password  // 登录得密码
	sqlStr := `SELECT user_id, username, password FROM user WHERE username=?`

	// user更新了里面得密码为数据库查询到得密码
	err = db.Get(user, sqlStr, user.Username)

	if err != nil {
		// 查询数据库失败
		return err
	}

	// 判断密码是否正确
	// 传进来得密码
	password := encryptPassword(oldPassword)

	if password != user.Password {
		return ErrorInvalidPassword
	}
	return
}
