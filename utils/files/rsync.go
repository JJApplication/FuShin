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
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path"
	"time"
)

// 调用rsync同步
// 使用同步的cmd调用

const (
	RsyncCMD = "rsync -q -p -backup %s %s"
	TarCMD   = "tar -zcf %s.tar.gz %s"
)

var TmpDir = os.TempDir()

// 同步的cmd调用
func cmdCall(cmd string) error {
	c := exec.Command("bash", "-c", cmd)
	return c.Run()
}

// 同步copy
// src源目录
// dst 目标目录 不存在自动创建
func Rsync(src, dst string) error {
	if dst == "" || dst == "/" {
		return errors.New("rsync dir is root path")
	}
	cmd := fmt.Sprintf(RsyncCMD, src, dst)
	return cmdCall(cmd)
}

// Tar 会进行目录切换 需要保证目录src存在
func Tar(file, src string) error {
	parentDir := path.Dir(src)
	if parentDir == "" || parentDir == "/" {
		return errors.New("tar dir is root path")
	}
	err := os.Chdir(parentDir)
	if err != nil {
		return err
	}
	baseName := path.Base(src)
	cmd := fmt.Sprintf(TarCMD, AutoName(file), baseName)
	return cmdCall(cmd)
}

// RsyncAndTar 自动备份并压缩
// eg: RsyncAndTar("/test", "/tmp/back", false)
// 会自动压缩目录到/tmp/back-timestamp.tar.gz
func RsyncAndTar(src, dst string, postClear bool) error {
	err := Rsync(src, dst)
	if err != nil {
		return err
	}
	err = Tar(dst, dst)
	if err != nil {
		return err
	}
	if postClear {
		return clear(dst)
	}
	return nil
}

// AutoName 自动创建时间戳目录
func AutoName(dst string) string {
	timeStamp := time.Now().Format("2006-01-02-15-04-05.00000")
	return fmt.Sprintf("%s-%s", dst, timeStamp)
}

// 清理目录
func clear(p string) error {
	if p == "" || p == "/" {
		return errors.New("clear dir is root path")
	}
	return os.RemoveAll(p)
}
