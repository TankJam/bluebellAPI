package middlewares

import (
	"bluebellAPI/controller"
	"bluebellAPI/pkg/jwt"
	"github.com/gin-gonic/gin"
	"strings"
)

// JWTAuthMiddleware JWT认证中间件
func JWTAuthMiddleware() gin.HandlerFunc{

	return func(c *gin.Context) {
		/*
			客户端携带Token的三种方式:
				1.放在请求头
				2.放在请求体
				3.放在URI中

			- 此版本在请求头中携带Token
		*/
		authHeader := c.Request.Header.Get("Authorization")  // 从请求头中获取token
		if authHeader == "" {
			controller.ResponseError(c, controller.CodeNeedLogin)  // 未登录，返回需要登录
			c.Abort()
			return
		}

		// 按空格分割
		parts := strings.SplitN(authHeader, " ", 2)
		// 若长度不等于2，以及第一个参数不是Bearer就返回Token错误
		// Bearer token值
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			controller.ResponseError(c, controller.CodeInvalidToken)
			c.Abort()
			return
		}

		// parts[1] 是获取到tokenString，我们使用之前定义好的解析JWT的函数来解析它
		mc, err := jwt.ParseToken(parts[1])
		if err != nil {
			controller.ResponseError(c, controller.CodeInvalidToken)
			c.Abort()
			return
		}

		// 将当前请求的userID信息保存到请求的上下文 c 上
		c.Set(controller.CtxUserIDKey, mc.UserID)
		c.Next()
	}
}
