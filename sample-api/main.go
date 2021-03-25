package main

///////////////////////
// 实现一个基础的RESTfulAPI服务器 //
///////////////////////
import (
	"errors"
	"log"
	"net/http"
	"sample-api/router"
	"time"

	"github.com/gin-gonic/gin"
)

///////////////////////////////////////////////////
// main()为入口函数，主要做一些配置文件解析、程序初始化和路由加载之类的事 //
// 最终调用 http.ListenAndServe() 在指定端口启动一个 HTTP 服务器 //
///////////////////////////////////////////////////
func main() {
	// 创建一个不包含中间件的Gin实例
	g := gin.New()

	// 定义容纳多个中间件（gin.HandlerFunc{}类型）的Slice
	middlewares := []gin.HandlerFunc{}

	// main() 函数通过调用 router.Load 函数来加载路由，函数路径为 router/router.go
	router.Load(
		// Cores.
		g,

		// 中间件，middlewares...表示为可变参数middlewares传参
		middlewares...,
	)

	// 通过访问 pingServer() 函数 验证 API Server 是否工作正常
	go func() {
		if err := pingServer(); err != nil {
			log.Fatal("The router has no response, or it might took too long to start up.", err)
		}
		log.Print("The router has been deployed successfully.")
	}()

	log.Printf("Start to listening the incoming requests on http address: %s", ":8080")
	// 调用 http.ListenAndServe() 在指定端口启动一个 HTTP 服务器，并输出错误信息
	log.Printf(http.ListenAndServe(":8080", g).Error())
}

// pingServer函数通过访问 /sd/health接口验证 API Server 是否工作正常
// pingServer函数向/sd/health发送GET请求，如果函数正确执行并且返回的 HTTP StatusCode 为 200，则说明 API 服务器可用，pingServer函数输出部署成功提示；如果超过指定次数，pingServer 直接终止 API Server 进程
func pingServer() error {
	for i := 0; i < 2; i++ {
		// 向 /sd/health 接口发送 GET 请求
		resp, err := http.Get("http://127.0.0.1:8080" + "/sd/health")
		if err == nil && resp.StatusCode == 200 {
			return nil
		}

		// Sleep 1秒后再次执行检查
		log.Print("Waiting for the router, retry in 1 second.")
		time.Sleep(time.Second)
	}
	return errors.New("Cannot connect to the router.")
}
