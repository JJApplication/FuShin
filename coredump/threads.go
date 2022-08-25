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

// threads

const (
	ThreadCore = ".threads_profile"
)

func GenerateThreadsCore(path string) error {
	var err error
	if path == "" {
		path = "."
	}
	if !files.IsExist(path) {
		path = DefaultDumpDir
	}
	cpuFile := filepath.Join(path, getExecName()+ThreadCore)
	f, err := os.Create(cpuFile)
	if err != nil {
		return err
	}
	err = pprof.Lookup("threadcreate").WriteTo(f, 0)
	f.Close()
	return err
}
