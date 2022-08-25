//go:build linux || darwin

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
	"syscall"

	"github.com/JJApplication/fushin/utils/files"
)

// 记录stderr到指定目录下

const (
	DefaultDumpDir = "/tmp"
	CoreDumpSuffix = ".coredump"
)

// CoreDump path为coredump文件存储路径 不存在时使用DefaultDumpDir
// 当标准日志被输出到stderr时也会被记录
func CoreDump(path string) {
	execName := getExecName()
	savePath := DefaultDumpDir
	if files.IsExist(path) {
		savePath = path
	}
	dumpFile := filepath.Join(savePath, execName+CoreDumpSuffix)
	logFile, err := os.OpenFile(dumpFile, os.O_CREATE|os.O_RDWR|os.O_SYNC|os.O_TRUNC, 0644)
	if err != nil {
		return
	}

	err = syscall.Dup3(int(logFile.Fd()), int(os.Stderr.Fd()), 0)
	if err != nil {
		return
	}
	return
}
