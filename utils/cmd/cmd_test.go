/*
Create: 2022/8/30
Project: FuShin
Github: https://github.com/landers1037
Copyright Renj
*/

// Package cmd
package cmd

import (
	"testing"
)

func TestCmder(t *testing.T) {
	res, err := Cmd.ExecStr("ls")
	t.Log(res)
	t.Log(err)
}
