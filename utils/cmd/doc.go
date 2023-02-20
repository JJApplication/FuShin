/*
Create: 2023/2/21
Project: FuShin
Github: https://github.com/landers1037
Copyright Renj
*/

// Package cmd 系统调用exec包
//
// 实现了命令的快速调用 使用New()创建一个Cmder对象
// c := NewCmder()
// c.Exec("ls")
//
// 实现了基于windows/linux的daemon模式启动
// cmd.Daemon() 或 cmd.DaemonCall(cmd)调用任意命令
package cmd
