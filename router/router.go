package router

import (
	"bluebellAPI/controller"
	"bluebellAPI/logger"
	"bluebellAPI/middlewares"
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
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	// 3.创建 api v1 路由组
	v1 := r.Group("/api/v1")

	// 4.注册
	v1.POST("/signup", controller.SignUpHandler)

	// 5.登录
	v1.POST("/login", controller.LoginHandler)

	// 6.加载 JWT 认证中间件
	v1.Use(middlewares.JWTAuthMiddleware())  // 放在登录注册之后

	// 7.主页、根据id查询、提交创建文章
	{
		// 社区主页
		v1.GET("/community", controller.CommunityHandler)
		// 根据id返回社区
		// http://127.0.0.1:9527/api/v1/community/1
		v1.GET("/community/:id", controller.CommunityDetailHandler)

		// 新增帖子
		v1.POST("/post", controller.CreatePostHandler)
		// 查询所有帖子 V1
		v1.GET("/posts", controller.GetPostListHandler)
		// 根据帖子id查询帖子
		v1.GET("/post/:id", controller.GetPostDetailHandler)
		// 根据时间查询所有帖子 V2
		v1.GET("/posts2", controller.GetPostListHandler2)

		// 投票功能
		v1.POST("/vote", controller.PostVoteHandler)

	}

	// last: 若路由错误，则返回 404
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})
	return r
}
