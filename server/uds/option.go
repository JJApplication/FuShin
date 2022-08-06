/*
Create: 2022/8/5
Project: FuShin
Github: https://github.com/landers1037
Copyright Renj
*/

// Package uds
package uds

type Option struct {
	AutoCheck         bool // 定时检查连接
	AutoCheckDuration int  // 单位s
	AutoRecover       bool // 出错后自动恢复
	RequestFormat     Req  // 请求的格式
	ResponseFormat    Res  // 响应的格式
	MaxSize           int  // 接收的byte最大字节数 默认1024
	LogTrace          bool // 是否使用内置的日志记录
}

var DefaultOption = Option{
	AutoCheck:         false,
	AutoCheckDuration: 60,
	AutoRecover:       false,
	ResponseFormat:    Res{},
	RequestFormat:     Req{},
	MaxSize:           1 << 10,
	LogTrace:          true,
}

var BigBodyOption = Option{
	AutoCheck:         false,
	AutoCheckDuration: 60,
	AutoRecover:       false,
	ResponseFormat:    Res{},
	RequestFormat:     Req{},
	MaxSize:           1 << 20,
}

var AutoCheckOption = Option{
	AutoCheck:         true,
	AutoCheckDuration: 60,
	AutoRecover:       false,
	ResponseFormat:    Res{},
	RequestFormat:     Req{},
	MaxSize:           1 << 10,
	LogTrace:          true,
}
