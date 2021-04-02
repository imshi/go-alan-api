package config

import (
	"log"
	"strings"
	"time"

	"github.com/fsnotify/fsnotify"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// 定义一个名为 Config 的结构体，并附加一些方法（解析配置、热更新），用来声明结构体变量存储参数的键值对
type Config struct {
	Name string
}

// config 包主要的功能函数，用以初始化配置
func Init(cfg string) error {
	c := Config{
		Name: cfg,
	}

	// 解析配置
	if err := c.initConfig(); err != nil {
		return err
	}

	// 初始化日志包，因为要读取日志相关的配置，所以需放在配置初始化函数 c.initConfig() 之后
	c.initLog()
	// 监控配置文件变化并热加载程序
	c.watchConfig()

	return nil
}

// 结构体方法一：解析配置文件
func (c *Config) initConfig() error {
	if c.Name != "" {
		viper.SetConfigFile(c.Name) // 如果指定了配置文件（ *cfg 不为空），则解析指定的配置文件
	} else {
		viper.AddConfigPath("conf") // 如果没有指定配置文件，则解析默认的配置文件（conf/config.yaml）
		viper.SetConfigName("config")
	}
	viper.SetConfigType("yaml") // 设置配置文件格式，这里为YAML
	viper.AutomaticEnv()        // 读取匹配的环境变量
	// 环境变量前缀不区分大小写,vipper会将小写前缀自动转化为大写再查找一遍
	viper.SetEnvPrefix("MYAPP")               // 读取环境变量的前缀为MYAPP
	replacer := strings.NewReplacer(".", "_") // 参数中的 . 对应环境变量中的 _
	viper.SetEnvKeyReplacer(replacer)
	if err := viper.ReadInConfig(); err != nil { // viper.ReadInConfig() 函数最终会调用 Viper 解析配置文件
		return err
	}

	return nil
}

// // 结构体方法二：设置日志输出格式 - 日志文件不分割的写法
// func (c *Config) initLog() {
// 	// 默认情况下，日志输出到io.Stderr。可以调用logrus.SetOutput传入一个io.Writer参数，后续调用相关方法日志将写到io.Writer中，支持定义多个writer，使用io.MultiWriter， 同时将日志写到bytes.Buffer、标准输出和文件中
// 	// 这里只输出日志到文本文件
// 	writerFile, err := os.OpenFile("log-record.log", os.O_WRONLY|os.O_CREATE, 0755)
// 	if err != nil {
// 		log.Fatalf("create log file failed:%v", err)
// 	}
// 	// 在输出日志中添加文件名和方法信息
// 	logrus.SetReportCaller(true)
// 	// 设置日志为 json 格式（默认为文本格式）
// 	logrus.SetFormatter(&logrus.JSONFormatter{})
// 	// 输出日志到文本文件
// 	logrus.SetOutput(io.Writer(writerFile))
// }

// 结构体方法二：设置日志输出格式 - 按日期分割日志文件
func (c *Config) initLog() {
	content, err := rotatelogs.New("server.log"+"-%Y%m%d%H%M",
		rotatelogs.WithLinkName("server.log"), // 生成软链，指向最新日志文件
		// WithMaxAge 和 WithRotationCount 两者不能同时设置
		rotatelogs.WithMaxAge(6*time.Minute), //日志保留时长：最小单位为分钟
		// WithRotationCount(5),        // 日志保留份数：默认7份，大于7份或到了清理时间触发清理
		rotatelogs.WithRotationTime(time.Minute), //轮转周期：默认1分钟（最小1分钟）分隔一次；
	)
	if err != nil {
		log.Printf("failed to create rotatelogs: %s", err)
		return
	}
	// 在输出日志中添加文件名和方法信息
	logrus.SetReportCaller(true)
	// 设置日志为 json 格式（默认为文本格式）
	logrus.SetFormatter(&logrus.JSONFormatter{})
	// 输出日志到文本文件
	logrus.SetOutput(content)
	// 设置打印日志的最低等级，所有等级：Trace、Debug、Info、Warn、Error、Fatal、Panic
	logrus.SetLevel(logrus.TraceLevel)
}

// 结构体方法三：监控配置文件变化并热加载程序
func (c *Config) watchConfig() {
	// 通过 fsnotify.Event 获取配置文件变动事件，viper.watchConfig 函数进行热更新
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		// 配置文件发生变更之后会调用的回调函数
		// log.Printf("Config file changed: %s", e.Name)
		// log.Infof("Config file changed: %s", e.Name)
		logrus.Infof("Config file changed: %s", e.Name)
	})
}
