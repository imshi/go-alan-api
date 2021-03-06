#### 通过 viper 包进行配置文件读取
- Viper具有如下特性:
    - 设置默认值
    - 可以读取如下格式的配置文件：JSON、TOML、YAML、HCL
    - 监控配置文件改动，并热加载配置文件
    - 从环境变量读取配置
    - 从远程配置中心读取配置（etcd/consul），并监控变动
    - 从命令行 flag 读取配置
    - 从缓存中读取配置
    - 支持直接设置配置项的值
- Viper 配置读取顺序：
    - viper.Set() 所设置的值
    - 命令行 flag
    - 环境变量
    - 配置文件
    - 配置中心：etcd/consul
    - 默认值

- 测试验证
    - 编译运行configuration-read后查看控制台输出，改动配置后重新运行查看控制台输出
