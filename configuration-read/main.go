package main

import (
	"configuration-read/config"
	"configuration-read/router"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// 声明一个名为config、短标志为c，默认值为空，描述为“configuration-read config file path.” 的选项标志
var (
	// 该函数返回一个字符串指针
	cfg = pflag.StringP("config", "c", "", "configuration-read config file path.")
)

//
func main() {
	// 对标志进行命令行解析
	pflag.Parse()

	// 初始化来自 config 私有包的配置
	// cfg 变量值从命令行 flag 传入，可以传值，比如 ./configuration-read -c config.yaml；也可以为空，如果为空会默认读取 conf/config.yaml
	if err := config.Init(*cfg); err != nil {
		panic(err)
	}

	// gin 框架提供了三种模式（debug、release、test）模式分别用于线下线上不同场景，此处从配置文件中获取 runmode
	gin.SetMode(viper.GetString("runmode"))

	// Create the Gin engine.
	g := gin.New()

	// 定义容纳多个中间件（gin.HandlerFunc{}类型）的Slice
	middlewares := []gin.HandlerFunc{}

	// Routes.
	router.Load(
		// Cores.
		g,

		// Middlwares.
		middlewares...,
	)

	// Ping the server to make sure the router is working.
	go func() {
		if err := pingServer(); err != nil {
			log.Fatal("The router has no response, or it might took too long to start up.", err)
		}
		log.Print("The router has been deployed successfully.")
	}()

	// 通过 viper 读取配置文件中设置的端口（addr）
	log.Printf("Start to listening the incoming requests on http address: %s", viper.GetString("addr"))
	log.Print(http.ListenAndServe(viper.GetString("addr"), g).Error())
}

// pingServer pings the http server to make sure the router is working.
func pingServer() error {
	for i := 0; i < viper.GetInt("max_ping_count"); i++ {
		// Ping the server by sending a GET request to `/health`.
		resp, err := http.Get(viper.GetString("url") + "/sd/health")
		if err == nil && resp.StatusCode == 200 {
			return nil
		}

		// Sleep for a second to continue the next ping.
		log.Print("Waiting for the router, retry in 1 second.")
		time.Sleep(time.Second)
	}
	return errors.New("Cannot connect to the router.")
}
