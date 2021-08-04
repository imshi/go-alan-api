#### 通过 JWT 给 API 添加认证功能 - 010
- 介绍如何通过 JWT 进行 API 身份验证
- 验证：
    - 用户登录：curl -XPOST -H "Content-Type: application/json" http://127.0.0.1:8080/login -d'{"username":"kong","password":"kong123"}'
    - 如果请求时不携带签发的 token，会禁止请求：curl -XPOST -H "Content-Type: application/json" http://127.0.0.1:8080/v1/user -d'{"username":"kong","password":"kong123"}'
    - 请求时携带 token：curl -XPOST -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.xxx.LjxrK9DuAwAzUD8-9v43NzWBN7HXsSLfebw92DKd1JQ" -H "Content-Type: application/json" http://127.0.0.1:8080/v1/user -d'{"username":"kong","password":"kong123"}'