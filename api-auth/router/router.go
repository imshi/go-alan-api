package router

import (
	"net/http"

	"api-auth/handler/sd"
	"api-auth/handler/user"
	"api-auth/router/middleware"

	"github.com/gin-gonic/gin"
)

// Load loads the middlewares, routes, handlers.
// Load函数导入中间件、路由、处理器
func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	// Middlewares.
	// 应用全局中间件为每一个请求设置Header
	g.Use(gin.Recovery())     //在处理某些请求时可能因为程序 bug 或者其他异常情况导致程序 panic，这时候为了不影响下一次请求的调用，需要通过 gin.Recovery()来恢复 API 服务器
	g.Use(middleware.NoCache) //强制浏览器不使用缓存，在router/middleware中实现
	g.Use(middleware.Options) //浏览器跨域 OPTIONS 请求设置，在router/middleware中实现
	g.Use(middleware.Secure)  //一些安全设置，在router/middleware中实现
	g.Use(mw...)
	// 404 Handler.
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API route.")
	})

	// 登录认证功能接口
	g.POST("/login", user.Login)
	// 用户操作的路由配置
	u := g.Group("/v1/user")
	u.Use(middleware.AuthMiddleware())
	{
		u.POST("", user.Create)       // 创建用户
		u.DELETE("/:id", user.Delete) // 删除用户
		u.PUT(":id", user.Update)     // 更新用户
		u.GET("", user.List)          // 创获取用户列表
		u.GET("/:username", user.Get) // 获取指定用户的详细信息
	}

	// The health check handlers
	// 定义了一个叫 sd 的路由组，在该分组下注册了 /health、/disk、/cpu、/ram HTTP 路径，分别路由到 sd.HealthCheck、sd.DiskCheck、sd.CPUCheck、sd.RAMCheck 函数，用来检查 API Server 的状态：健康状况、服务器硬盘、CPU 和内存使用量
	svcd := g.Group("/sd")
	{
		svcd.GET("/health", sd.HealthCheck)
		svcd.GET("/disk", sd.DiskCheck)
		svcd.GET("/cpu", sd.CPUCheck)
		svcd.GET("/ram", sd.RAMCheck)
	}

	return g
}
