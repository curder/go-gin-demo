# Go Gin Demo

## 下载

- 下载项目源文件
```
git clone git@github.com:curder/go-gin-demo.git
```

- 下载项目依赖
```
go mod download
```


## 配置

配置文件保存在`configs/app.yml`

- `server.port` 默认http服务器端口
- `database.driverName` 数据库驱动，默认使用MySQL数据库
  - `host` 数据库连接主机
  - `port` 数据库端口
  - `database` 数据库名
  - `user` 数据库连接用户名
  - `password` 数据库密码
  - `charset` 数据库字符集

## 编译

```
go build *.go
```

## 开发

```
go run main.go
```


## Links

- [Go-gin-vue-demo - 前端界面支持](https://github.com/curder/go-gin-vue-demo)
- [Go语言 Gin+Vue 前后端分离实战 - OceanLearn - Bilibili](https://www.bilibili.com/video/BV1CE411H7bQ)
- [Gin Framework](https://github.com/gin-gonic/gin)
