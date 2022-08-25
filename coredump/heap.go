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

// mem heap

const (
	HeapCore = ".heap_profile"
)

func GenerateHeapCore(path string) error {
	var err error
	if path == "" {
		path = "."
	}
	if !files.IsExist(path) {
		path = DefaultDumpDir
	}
	cpuFile := filepath.Join(path, getExecName()+HeapCore)
	f, err := os.Create(cpuFile)
	if err != nil {
		return err
	}
	err = pprof.WriteHeapProfile(f)
	f.Close()
	return err
}
