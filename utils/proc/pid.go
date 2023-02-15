/*
Create: 2023/2/15
Project: FuShin
Github: https://github.com/landers1037
Copyright Renj
*/

package proc

import (
	"io/ioutil"
	"os"
	"strconv"
)

// 获取运行时pid

func PID() int {
	return os.Getpid()
}

func PIDToFile(f string) error {
	return ioutil.WriteFile(f, []byte(strconv.Itoa(PID())), 0644)
}

func PPID() int {
	return os.Getppid()
}
