//go:build linux || darwin

/*
Create: 2023/2/14
Project: FuShin
Github: https://github.com/landers1037
Copyright Renj
*/

package cmd

import "syscall"

// Fork on UNIX
func Fork(arg0 string, argv []string, attr *syscall.ProcAttr) (int, error) {
	return syscall.ForkExec(arg0, argv, attr)
}
