/*
Create: 2022/7/7
Project: FuShin
Github: https://github.com/landers1037
Copyright Renj
*/

// Package private
package private

import "github.com/JJApplication/fushin/inner"

// 创建app 当前支持web服务 client服务 noengine代理服务 空模板服务
// 空模板只存在.octupus元数据

var (
	APPType        = []string{"web", "client", "noengine", "empty", "default"}
	APPTypeDefault = "empty"
	APPPrefix      = inner.COPYRIGHT
)
