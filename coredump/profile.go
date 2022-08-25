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
	"runtime/pprof"

	"github.com/JJApplication/fushin/utils/files"
)

// cpu profile

// 默认记录路径/tmp/*.cpu_profile

const (
	CpuCore = ".cpu_profile"
)

// GenerateCpuCore path为存储路径
// 不存在时使用默认路径/tmp
func GenerateCpuCore(path string) error {
	var err error
	if path == "" {
		path = "."
	}
	if !files.IsExist(path) {
		path = DefaultDumpDir
	}
	cpuFile := filepath.Join(path, getExecName()+CpuCore)
	cf, err := os.Create(cpuFile)
	if err != nil {
		return err
	}
	err = pprof.StartCPUProfile(cf)
	defer pprof.StopCPUProfile()
	cf.Close()
	return err
}
