/*
Create: 2022/8/25
Project: FuShin
Github: https://github.com/landers1037
Copyright Renj
*/

// Package coredump
package coredump

import (
	"testing"
)

func TestExecName(t *testing.T) {
	t.Log(getExecName())
}

func TestCore(t *testing.T) {
	p := ""
	t.Log(GenerateCpuCore(p))
	t.Log(GenerateHeapCore(p))
	t.Log(GenerateBlockCore(p))
	t.Log(GenerateThreadsCore(p))
	t.Log(GenerateGrsCore(p))
}
