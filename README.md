# goredis-admin
基于go语言开发的web版redis管理工具

## 简介
* 前端框架使用[element ui](https://github.com/ElemeFE/element)
* 服务端使用 [gin](https://github.com/gin-gonic/gin)


## 特性
* go语言编写，前后端分离，低功耗，高性能， 傻瓜式部署
* 支持管理多个redis服务
* 支持redis基本数据类型(string, list, hash, set, zset) 增、删、查、改
## 安装指南
1. 直接编译源码

   ```shell script
     git clone https://github.com/iv587/goredis-admin.git
     cd ./frontend
     npm install
     npm run build
     go mod download
     go mod vendor
     go build gredisw.go
   ```
2. 下载[二进制文件](https://github.com/iv587/goredis-admin/releases)

## 使用
默认用户名 admin ,密码：goredis
     
## 安装目录说明
```shell script
-- start                    可执行文件
-- data
  -- user.db                用户存放文件
  -- redis.db               redis配置存放文件
-- conf
  -- app.conf               应用配置存放文件
-- log                      日志文件目录
```
## LICENSE
[Apache License](./LICENSE)
