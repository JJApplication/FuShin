/*
Create: 2022/8/26
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

const (
	GoroutineCore = ".goroutine_profile"
)

func GenerateGrsCore(path string) error {
	var err error
	if path == "" {
		path = "."
	}
	if !files.IsExist(path) {
		path = DefaultDumpDir
	}
	cpuFile := filepath.Join(path, getExecName()+GoroutineCore)
	f, err := os.Create(cpuFile)
	if err != nil {
		return err
	}
	err = pprof.Lookup("goroutine").WriteTo(f, 0)
	f.Close()
	return err
}
