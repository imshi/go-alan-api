#### 读取和返回 HTTP 请求
> 读取主要用 gin 框架自带的函数，返回要统一用函数 SendResponse
- 功能特性
    - 如何读取 HTTP 请求数据
    - 如何返回数据
    - 如何定制业务的返回格式
- 测试：
    - 发送 HTTP 请求查看返回数据：`curl -XPOST -H "Content-Type: application/json" http://127.0.0.1:8080/v1/user/admin2?desc=test -d'{"username":"admin","password":"admin"}'`
