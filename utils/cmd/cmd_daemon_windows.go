//go:build windows

/*
Create: 2023/2/15
Project: FuShin
Github: https://github.com/landers1037
Copyright Renj
*/

package cmd

import (
	"os"
	"os/exec"
)

// 以fork的形式启动进程

// DaemonCall daemon调用任意指定的命令
func DaemonCall(program string, args ...string) error {
	cmd := exec.Command(program, args...)
	if err := cmd.Start(); err != nil {
		return err
	}

	return nil
}

// Daemon 以daemon模式运行当前的程序
// 限制需要指定daemon arg来用于排除
func Daemon(exclude string) {
	procName := os.Args[0]
	var args []string
	for _, arg := range os.Args[1:] {
		if arg != exclude {
			args = append(args, arg)
		}
	}
	cmd := exec.Command(procName, args...)
	_ = cmd.Start()
}

// DaemonWith 返回pid和错误
func DaemonWith(exclude string) (int, error) {
	procName := os.Args[0]
	var args []string
	for _, arg := range os.Args[1:] {
		if arg != exclude {
			args = append(args, arg)
		}
	}
	cmd := exec.Command(procName, args...)

	err := cmd.Start()
	return cmd.Process.Pid, err
}
