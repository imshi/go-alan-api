#### 通过 gorm 连接 mysql 数据库
> 这里使用v2版本的gorm，包如下：

- gorm.io/gorm
- gorm.io/driver/mysql

> 额外补充说明：

1. 因为 v1 版本中使用体验不佳等问题，gorm v2版本中移除了用来关闭连接的Close，使用连接池设置可复用时间即可（gorm自行处理）
2. init.go文件汇总为了让 API 服务器可能需要同时访问多个数据库，定义了一个叫 Database 的 struct ，导致配置巨复杂，可酌情修改；
3. mysql8中设置字符集的变量更换为 character_set_server ，创建数据库连接的时候将 charset 替换为 character_set_server；
