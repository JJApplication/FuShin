/*
Create: 2022/8/5
Project: FuShin
Github: https://github.com/landers1037
Copyright Renj
*/

// Package uds
package uds

import (
	"fmt"
	"net"
	"os"
	"strings"
	"syscall"

	"github.com/JJApplication/fushin/utils/json"
)

// RemoveSocket 删除socket
func RemoveSocket(s string) error {
	if strings.HasSuffix(s, ".sock") {
		return syscall.Unlink(s)
	}
	return syscall.Unlink(fmt.Sprintf("%s.sock", s))
}

// CreateSocket 创建socket 默认监听时会创建遇到无权限时调用此方法创建
// s app or app.socket
func CreateSocket(s string) error {
	_ = RemoveSocket(s)
	if strings.HasSuffix(s, ".sock") {
		_, err := os.Create(s)
		return err
	}
	f, err := os.Create(fmt.Sprintf("%s.sock", s))
	_ = f.Close()
	return err
}

// GetSocket 获取sock文件每次使用前清空之前的连接
func GetSocket(s string) string {
	_ = RemoveSocket(s)
	if strings.HasSuffix(s, ".sock") {
		return s
	}

	return fmt.Sprintf("%s.sock", s)
}

// Response 响应body数据
func Response(c net.Conn, res Res) error {
	data, err := json.Json.Marshal(res)
	if err != nil {
		return err
	}
	_, err = c.Write(data)
	return err
}

// ModuleName 模块名称
func ModuleName() string {
	return moduleName
}
