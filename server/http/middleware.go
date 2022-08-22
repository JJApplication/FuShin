/*
Create: 2022/8/10
Project: FuShin
Github: https://github.com/landers1037
Copyright Renj
*/

// Package http
package http

// 内置中间件

// MiddleWareLogger 路由日志
func MiddleWareLogger() WrapperFunc {
	return convertHandle(ginLogger())
}

// MiddleWareCors 默认的cors
// 允许所有跨域
func MiddleWareCors() WrapperFunc {
	return convertHandle(ginCors())
}
