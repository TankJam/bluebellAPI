package router

import (
	"bluebellAPI/controller"
	"github.com/gin-gonic/gin"
	"net/http"
)

// SetupRouter 注册路由函数
func SetupRouter(mode string) *gin.Engine {
	// 判断mode
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode) // gin设置成发布模式
	}

	// 1.实例化并获取gin引擎对象
	r := gin.New()

	// 2.加载 日志 中间件
	//r.Use()

	// 3.创建 api v1 路由组
	v1 := r.Group("/api/v1")

	// 4.注册
	v1.POST("/signup", controller.SignUpHandler)

	// 5.登录

	// 6.加载 JWT 认证中间件

	// 7.主页、根据id查询、提交创建文章

	// last: 若路由错误，则返回 404
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})
	return r
}
