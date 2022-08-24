/*
Create: 2022/8/24
Project: FuShin
Github: https://github.com/landers1037
Copyright Renj
*/

// Package cmd
package cmd

import (
	"errors"
	"os"
	"strings"

	"github.com/JJApplication/fushin/utils/files"
)

const (
	BASH      = ".sh"
	ShellPerm = 0700
	EmptyBash = "#!/usr/bin/env bash"
)

// NewEmptyBash 新建一个空shell文件
// dst为文件生成的目录
// dst为.sh结尾的文件
func NewEmptyBash(dst string) error {
	// file
	if strings.HasSuffix(dst, BASH) {
		if files.IsExist(dst) {
			return errors.New("bash file already exist")
		}
		f, err := os.OpenFile(dst, os.O_CREATE|os.O_RDWR, ShellPerm)
		if err != nil {
			return err
		}
		_, err = f.WriteString(EmptyBash)
		return err
	}
	if !files.IsExist(dst) {
		return errors.New("dst path is not exist")
	}
	return nil
}

// NewBash 新建一个带命令的shell文件
func NewBash(dst, sh string) error {
	if strings.HasSuffix(dst, BASH) {
		if files.IsExist(dst) {
			return errors.New("bash file already exist")
		}
		f, err := os.OpenFile(dst, os.O_CREATE|os.O_RDWR, ShellPerm)
		if err != nil {
			return err
		}
		unixSh := strings.ReplaceAll(sh, "\r\n", "\n")
		_, err = f.WriteString(unixSh)
		return err
	}
	if !files.IsExist(dst) {
		return errors.New("dst path is not exist")
	}
	return nil
}
