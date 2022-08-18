//go:build linux || darwin

/*
Create: 2022/8/18
Project: FuShin
Github: https://github.com/landers1037
Copyright Renj
*/

// Package files
package files

import (
	"testing"
)

func TestRsyncAndTar(t *testing.T) {
	err := RsyncAndTar("/home/test", "/tmp/1", true)
	if err != nil {
		t.Log(err)
	}
}
