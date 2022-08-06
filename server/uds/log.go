/*
Create: 2022/8/5
Project: FuShin
Github: https://github.com/landers1037
Copyright Renj
*/

// Package uds
package uds

// Logger 自定义的日志记录器 实现接口即可
type Logger interface {
	Info(v ...interface{})
	Warn(v ...interface{})
	Error(v ...interface{})
	InfoF(fmt string, v ...interface{})
	WarnF(fmt string, v ...interface{})
	ErrorF(fmt string, v ...interface{})
}
