# FuShin
FuShin Stone

`Fushin`是一个适用于JJApplication的一站式微服务模板生成库

使用`fushin` 生成对应的微服务模板

或者使用`go module`Fushin搭建一个JJApplication

## 微服务类型
在JJApplication中支持以下服务类型
- web
- module
- noengine
- client
- empty

`web`为典型的提供web体验的微服务
`module`为服务的模块
`noengine`为典型的NoEngine服务
`client`为典型的客户端服务
`empty`为一个空白的模板

## 微服务通信协议
默认支持的通信协议在对用户侧基于http, http2, ws
在服务端基于http或tcp的rpc, 基于unix domain的uds, 和自定义`protof - Plnack`

## 服务管理机制
所有的服务会生成一个模型文件`${App}.pig`
模型文件中记录了服务的基础信息，通信协议，对外端口等
模型文件由`octopusTree`统一管理，由`Apollo`服务解析存储，并做统一的管理调度

## 服务间通信机制
在服务器上的各个微服务之间的通信全部基于`unix domain`
由服务`octopusTwig`管理调用服务间uds的通信传输

## 服务模型设计
当前模型文件为基于`json`的`pig`文件, 支持环境变量和自定义字段覆盖
后期支持`yml`的文件格式
模型文件的定义基于`octopus_meta`