/*
Create: 2023/3/27
Project: FuShin
Github: https://github.com/landers1037
Copyright Renj
*/

// Package autocert
//
// 自动生成更新ssl证书 基于acme
// 简单的http认证方式，可以启动在:http以供challenge时使用
//
// 使用
// cert := autocert.NewDefault()
//
// cert.RunAndStop()
// 证书将会保存至fushin-cert目录下
package autocert
