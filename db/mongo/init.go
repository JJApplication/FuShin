/*
Create: 2022/8/17
Project: FuShin
Github: https://github.com/landers1037
Copyright Renj

mongodb的客户端
基于mgm实现一个全局的单例客户端
*/

// Package mongo
package mongo

const (
	moduleName = "<MongoClient>"
)

// ModuleName 模块名称
func ModuleName() string {
	return moduleName
}
