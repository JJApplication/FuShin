/*
Create: 2023/2/21
Project: FuShin
Github: https://github.com/landers1037
Copyright Renj
*/

// Package env env包提供环境变量的操作同时支持dotenv
//
// 使用NewEnvLoader()创建一个环境加载器实例
// 获取初始类型string的环境变量
//
// e := NewEnvLoader()
// e.Get("KEY").Raw()
// 自动转换为Int类型
// e.Get("KEY").Int()
//
// 使用dotenv来加载文件，默认会读取.env配置文件
// env.Load() // 传入file1 file2
//
// 支持以开发模式的方式自动加载，根据环境变量GO_DEVMODE配置
// env.AutoLoad()
package env
