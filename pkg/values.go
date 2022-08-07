/*
Create: 2022/6/29
Project: FuShin
Github: https://github.com/landers1037
Copyright Renj
*/

// Package pkg 内部需要使用的常量
package pkg

const (
	Fushin = "Fushin"
)

const (
	FushinEnableInitLog  = "FushinEnableInitLog"
	FushinEnableDebugLog = "FushinEnableDebugLog"
)

var (
	FushinMode     = "development"
	FushinLogColor = "true"
	FushinPanic    = "false"
)

const (
	dev = iota
	prod
	blank
	fatal
)

// SetFushinMode 设置fushin的全局变量
// mode:
// dev: 0
// prod: 1
// blank: 2 禁止彩色输出
// fatal: 3 在出错时panic退出
func SetFushinMode(mode int) {
	switch mode {
	case dev:
		FushinMode = "development"
		FushinLogColor = "true"
	case prod:
		FushinMode = "production"
		FushinLogColor = "true"
	case blank:
		FushinLogColor = "false"
	case fatal:
		FushinPanic = "true"
	default:
		FushinMode = "development"
		FushinLogColor = "true"
	}
}
