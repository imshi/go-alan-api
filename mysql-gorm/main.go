package main

///////////////////////
// 实现一个基础的RESTfulAPI服务器 //
///////////////////////
import (
	"errors"
	"mysql-gorm/config"
	"mysql-gorm/model"
	"mysql-gorm/router"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// 声明一个名为config、短标志为c，默认值为空，描述为“config file path.” 的选项标志，用来捕获存储命令行选项参数；
var (
	// 该函数返回一个字符串指针
	cfg = pflag.StringP("config", "c", "", "config file path.")
)

///////////////////////////////////////////////////
// main()为入口函数，主要做一些配置文件解析、程序初始化和路由加载之类的事 //
// 最终调用 http.ListenAndServe() 在指定端口启动一个 HTTP 服务器 //
///////////////////////////////////////////////////
func main() {
	// 对标志进行命令行解析
	pflag.Parse()

	// 初始化来自 config 私有包的配置
	// cfg 变量值从命令行 flag 传入，可以传值，比如 ./mysql-gorm -c config.yaml；也可以为空，如果为空会默认读取 conf/config.yaml
	if err := config.Init(*cfg); err != nil {
		panic(err)
	}

	model.DB.Init()
	// 验证热更新：修改配置文件中 runmode 的值，观察控制台输出（该 for 循环会阻塞程序向下运行，验证完毕及时剔除，否则 for 循环以下代码不会执行）
	// for {
	// 	fmt.Println(viper.GetString("runmode"))
	// 	time.Sleep(4 * time.Second)
	// }

	// gin 框架提供了三种模式（debug、release、test）模式分别用于线下线上不同场景，此处从配置文件中获取 runmode
	gin.SetMode(viper.GetString("runmode"))

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
			logrus.Fatal("The router has no response, or it might took too long to start up.", err)
		}
		logrus.Info("The router has been deployed successfully.")
	}()

	// 通过 viper 读取配置文件中设置的端口（addr）
	logrus.Infof("Start to listening the incoming requests on http address: %s", viper.GetString("addr"))
	logrus.Info(http.ListenAndServe(viper.GetString("addr"), g).Error())
}

// pingServer函数通过访问 /sd/health接口验证 API Server 是否工作正常
// 如果函数正确执行并且返回的 HTTP StatusCode 为 200，则说明 API 服务器可用，pingServer函数输出部署成功提示；如果超过指定次数，pingServer 直接终止 API Server 进程
func pingServer() error {
	for i := 0; i < viper.GetInt("max_ping_count"); i++ {
		// 向 /sd/health 接口发送 GET 请求
		resp, err := http.Get(viper.GetString("url") + "/sd/health")
		if err == nil && resp.StatusCode == 200 {
			return nil
		}

		// Sleep 1秒后再次执行检查
		logrus.Info("Waiting for the router, retry in 1 second.")
		time.Sleep(time.Second)
	}
	return errors.New("Cannot connect to the router.")
}
