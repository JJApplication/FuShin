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
	BlockCore = ".block_profile"
)

func GenerateBlockCore(path string) error {
	var err error
	if path == "" {
		path = "."
	}
	if !files.IsExist(path) {
		path = DefaultDumpDir
	}
	cpuFile := filepath.Join(path, getExecName()+BlockCore)
	f, err := os.Create(cpuFile)
	if err != nil {
		return err
	}
	err = pprof.Lookup("block").WriteTo(f, 0)
	f.Close()
	return err
}
