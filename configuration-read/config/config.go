package config

import (
	"log"
	"strings"

	"github.com/fsnotify/fsnotify"
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

// 结构体方法二：监控配置文件变化并热加载程序
func (c *Config) watchConfig() {
	// 通过 fsnotify.Event 获取配置文件变动事件，viper.watchConfig 函数进行热更新
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		// 配置文件发生变更之后会调用的回调函数
		log.Printf("Config file changed: %s", e.Name)
	})
}
