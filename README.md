<!-- TOC -->
* [DATA KNIFE](#data-knife)
  * [简介](#简介)
  * [开发相关](#开发相关)
  * [部署指南](#部署指南)
    * [源码编译部署](#源码编译部署)
    * [Docker容器部署](#docker容器部署)
    * [二进制安装](#二进制安装)
  * [开发里程碑](#开发里程碑)
    * [当前版本实现的功能:](#当前版本实现的功能-)
    * [项目愿景](#项目愿景)
  * [LICENSE](#license)
<!-- TOC -->

# DATA KNIFE
> 一款简单易用的web数据管理工具

## 简介
> * 使用Go语言开发，低功耗，高性能，
> * 前后端端分离
> * 傻瓜式一键部署
> * 支持多个数据实例的操作管理

## 开发相关
* go1.18 + vue3
* 前端框架使用[element plus](https://github.com/element-plus/element-plus) + vite + typescript
* 服务端使用 [gin](https://github.com/gin-gonic/gin), [gorm](https://gorm.io/gorm), [redigo](https://github.com/gomodule/redigo) 等

## 部署指南

### 源码编译部署
```bash
# 编译页面 并将打包文件输出到web目录中
cd ./data-knife-ui
yarn
yarn build

# 编译后端
go mod download
go mod vendor
go build DataKnife.go

# 以8080端口启动
./DataKnife
```
### Docker容器部署

### 二进制安装
 
## 开发里程碑
### 当前版本实现的功能:
> * 实现redis的string，hash数据操作功能
> * 实现redis的info面板显示工呢

### 项目愿景
> 完善redis数据操作类型
> 增加mysql、mongo、etcd等数据界面管理


## LICENSE
[Apache License](./LICENSE)
