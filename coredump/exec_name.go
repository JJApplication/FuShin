/*
Create: 2022/8/25
Project: FuShin
Github: https://github.com/landers1037
Copyright Renj
*/

// Package coredump
package coredump

import (
	"os"
	"path/filepath"
	"strings"
)

// 获取运行时文件名称 存在空格时转换
func getExecName() string {
	name := os.Args[0]
	name = filepath.Base(name)
	if len(strings.Fields(name)) > 1 {
		return strings.Join(strings.Fields(name), "_")
	}
	return name
}
