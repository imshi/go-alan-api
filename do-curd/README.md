#### 用户业务逻辑处理（增删改查）- 008
- 各种场景的业务逻辑处理
    - 创建、删除、查询、更新用户；查询用户列表；查询用户指定信息
- 数据库CURD操作
    - 创建用户：`curl -XPOST -H "Content-Type: application/json" http://127.0.0.1:8080/v1/user -d'{"username":"kong","password":"kong123"}'`
    - 查询用户列表：`curl -XGET -H "Content-Type: application/json" http://127.0.0.1:8080/v1/user -d'{"offset": 0, "limit": 20}'`
    - 获取指定用户名的用户详细信息：`curl -XGET -H "Content-Type: application/json" http://127.0.0.1:8080/v1/user/kong`
    - 根据用户id更新指定用户信息：`curl -XPUT -H "Content-Type: application/json" http://127.0.0.1:8080/v1/user/2 -d'{"username":"kong","password":"kongmodify"}'`
    - 获取指定用户名的用户详细信息（验证改动结果）：`curl -XGET -H "Content-Type: application/json" http://127.0.0.1:8080/v1/user/kong`
    - 根据用户id删除指定用户：`curl -XDELETE -H "Content-Type: application/json" http://127.0.0.1:8080/v1/user/2`
    - 获取用户列表（分页，取出前20条数据）：`curl -XGET -H "Content-Type: application/json" http://127.0.0.1:8080/v1/user -d'{"offset": 0, "limit": 20}'`

- Gorm v2中，Save（..）将所有字段写入数据库；如果更新时没有传递，将会以零值写入数据库；本节中在调用更新用户操作时会改动createdAt字段并赋值为空（暂时不知道为什么），在MySQL8.0以上版本中不支持零日期格式，导致gorm插入默认数据出错，临时解决方法（此时 update操作将会成功执行，但是 createdAt 字段将被置为null）：
    - 把 CreatedAt 字段的日期类型time.Time改为指针类型*time.Time；
    - 设置 CreatedAt 字段允许为空