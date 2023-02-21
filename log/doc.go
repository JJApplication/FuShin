/*
Create: 2023/2/22
Project: FuShin
Github: https://github.com/landers1037
Copyright Renj
*/

// Package log 日志记录器
//
// 基于zap的日志库 提供zap可定制的日志类
//
// 初始化日志类
// logger = log.New()
// logger.Info()
// logger.Warn()
// logger.Error()
//
// 默认内置了常用的日志类
// prodLog = log.Default()
// devLog = log.Dev()
// 创建原生Logger而不是Sugar log.NewZap()
// 基于zapCore创建 zap.NewCore() 注意：带日志绕接的日志类只能通过core创建
package log
