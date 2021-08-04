#### HTTP 调用添加自定义处理逻辑 - 009
- 介绍如何用 gin middleware 特性给 API 添加唯一请求 ID
- 命令行执行curl命令获取用户列表，观察响应http头信息（部分接口调试工具显示的响应头信息不全）：
    - curl -v -XGET -H "Content-Type: application/json" http://127.0.0.1:8080/v1/user