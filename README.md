### golang企业级RESTful API实践记录
#### 部分说明：
- 使用 Modules 包管理器管理 packages 依赖，使用说明如下（以demo1为例）：

  - 提前配好`GOPROXY`环境变量，使用 GoCenter 加速依赖包下载
  - 复制仓库中的`demo01`至任意目录，进入该目录执行初始化：`go mod init demo01`，加载依赖：`go mod tidy`；
  - 整个目录全局修改`apiserver`为`demo1`（搜索替换）
  - 编译：`go build main.go`，运行编译结果(Windows下)：`main.exe`

---
#### 目录说明
- **sample-api**：实现一个基础的RESTfulAPI服务器
- **demo02**：使用Viper进行配置文件读取
