package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
	"time"
)

/*
- jwt实现
	MyClaims 自定义声明结构体并内嵌jwt.StandardClaims
	jwt包自带的jwt.StandardClaims只包含了官方字段
	我们这里需要额外记录一个username字段，所以要自定义结构体
	如果想要保存更多信息，都可以添加到这个结构体中
*/

// 1.定义一个盐
var mySecret = []byte("tank jam 是非常帅气的小伙~")


// MyClaims 2.定义一个基于 jwt 的 StandardClaims 派生出来的结构体
type MyClaims struct {
	UserID   int64  `json:"user_id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

// GenToken 生成JWT
func GenToken(userID int64, username string) (string, error) {
	// 创建一个我们自己的声明数据
	c := MyClaims{
		UserID:         userID,
		Username:       username,  // 自定义字段
		StandardClaims: jwt.StandardClaims{  // jwt的设置，比如过期时间，签发人等...
			// ExpiresAt: 设置过期时间
			ExpiresAt: time.Now().Add(
				// 获取yaml文件中的auth下的jwt有效时间
				time.Duration(viper.GetInt("auth.jwt_expire")) * time.Hour).Unix(),

			// Issuer: 签发人
			Issuer: "bluebell",
		},
	}

	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)

	// 使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString(mySecret)
}





