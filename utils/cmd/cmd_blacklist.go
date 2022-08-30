/*
Create: 2022/8/30
Project: FuShin
Github: https://github.com/landers1037
Copyright Renj
*/

// Package cmd
package cmd

import (
	"strings"
	"syscall"
)

// 不允许操作某些系统命令
// 支持设置umask并在后面继承此环境变量

var blackList = []string{"rm -rf /*", "reboot", "sleep", "hibernate", "shutdown", "halt", "init", "poweroff"}

func check(sh string) bool {
	for _, l := range blackList {
		if strings.Contains(sh, l) {
			return false
		}
	}
	return true
}

func umask(mask int) {
	if mask <= 0 {
		return
	}
	syscall.Umask(mask)
}
