### golang企业级RESTful API开发实践
#### 部分说明：
- 使用 Modules 包管理器管理 packages 依赖，使用说明如下（以sample-api为例）：

  - 提前配好`GOPROXY`环境变量，使用 goproxy.io 加速依赖包下载
  - 复制仓库中的`sample-api`至任意目录，进入该目录执行初始化：`go mod init sample-api`，加载依赖：`go mod tidy`；
  - 调试运行：go run main.go
  - 编译：`go build -v main.go`，制品位置(Windows下)：`sample-api.exe`；linux下：`sample-api`
  - 编译运行示例：
    - Windows下：双击 exe 文件
    - linux下：`./sample-api`
- 基于go 1.12及以上版本

---
#### 目录说明
1. **sample-api**：实现一个基础的 RESTfulAPI 服务器
2. **configuration-read**：使用 Viper 包进行配置文件读取
  - 调试运行示例(Linux)：`go run main.go -c /tmp/apidemo/config.yaml`
3. **log-record**：使用 lexkong/log 包进行日志输出格式定制
4. **mysql-gorm**：通过v2版本的 gorm 连接 mysql 数据库
5. **diy-errno**：自定义错误码，并且前后台展示不同内容
6. **do-request**：读取和返回 HTTP 请求
7. **do-curd**：用户业务逻辑处理（API核心功能和数据库增删改查）
8. **gin-middleware**：HTTP 调用添加自定义处理逻辑（添加唯一请求 ID）
9. **api-auth**：使用JWT进行 API 身份验证
