#### 启动一个基础的 RESTful API 服务器
- 功能特性：
    - 设置 HTTP Header
    - API 服务器健康检查和状态查询
- 测试验证（发送如下 HTTP GET 请求获取指定信息）：
    - `curl -XGET http://127.0.0.1:8080/sd/health`
    - `curl -XGET http://127.0.0.1:8080/sd/disk`
    - `curl -XGET http://127.0.0.1:8080/sd/cpu`
    - `curl -XGET http://127.0.0.1:8080/sd/ram`
