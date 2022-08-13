/*
Create: 2022/8/13
Project: FuShin
Github: https://github.com/landers1037
Copyright Renj
*/

// Package log
package log

// 日志接口
// 所有可扩展的日志均使用该接口以覆盖内部日志记录器

type LoggerInterface interface {
	Info(v ...interface{})
	Warn(v ...interface{})
	Error(v ...interface{})
	InfoF(fmt string, v ...interface{})
	WarnF(fmt string, v ...interface{})
	ErrorF(fmt string, v ...interface{})
}

type LoggerPanic interface {
	Panic(v ...interface{})
	PanicF(fmt string, v ...interface{})
}
