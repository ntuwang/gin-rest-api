# Gin-rest-api Web

## 介绍
一个极简http接口项目 基于gin+gorm+jwt+swagger的gin脚手架
- 路由
- 日志
- 数据库|ORM
- 配置管理
- 控制器
- 中间件
- 鉴权

主要包含以下API：

| METHOD | URI              |DESCRIPTION|
|--------|------------------|---|
| GET    | /                |默认首页
| POST   | /api/v1/auth     |登录认证
| GET    | /api/v1/tags     |标签列表
| POST   | /api/v1/tags     |发布标签
| PUT    | /api/v1/tags/:id |修改标签
| DELETE | /api/v1/tags/:id |删除标签

### 准备

创建一个 `gin-rest-api` 数据库，并且导入建表的 [databases](docs%2Fdatabases)

根据[app.ini.example](app.ini.example) 修改 `app.ini`创建配置文件

```
[app]
env = dev
http_port = 8080
#日志存储文件
log_file = ./app.log
#是否在控制台输出日志
log_console = true
log_level = debug # debug模式会开启db操作日志

#默认数据库
[db]
dialect = mysql
dsn = """root:123456@tcp(127.0.0.1:3306)/gin-rest-api?charset=utf8mb4&parseTime=True&loc=Local"""
max_idle_conn = 5
max_open_conn = 50

[redis]
host = 127.0.0.1
port = 6379
pass =
min_idle = 10
...
```

### 运行
```
$ go mod tidy
$ go run main.go 
```

Swagger 文档
[https://127.0.0.1:8080/swagger/index.html)