/*
Create: 2022/8/5
Project: FuShin
Github: https://github.com/landers1037
Copyright Renj

UDS Server与client搭配使用
通信使用基于unixTCP
报文为可以序列化为json的byte字节流
传输的数据不可以被json序列化时抛弃此次请求报文

请求的格式化
请求格式包含{Operation, Data}
operation用于区分在server中定义的报文处理函数
data为报文携带的数据 为json可序列的字符串
*/

// Package uds
package uds
