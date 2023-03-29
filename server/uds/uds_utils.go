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

	"github.com/JJApplication/fushin/utils/stream"
)

var udsResponseBuilder int = stream.JSONType

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
	data, err := buildStream(res)
	if err != nil {
		return err
	}
	_, err = c.Write(data)
	return err
}

// ResponseAny 响应任意数据
func ResponseAny(c net.Conn, res interface{}) error {
	data, err := buildStream(res)
	if err != nil {
		return err
	}
	_, err = c.Write(data)
	return err
}

// SetResponseBuilder 设置uds通信使用的流格式JSON YAML GOB FushinBuf
func SetResponseBuilder(buildType int) {
	udsResponseBuilder = buildType
}

func buildStream(v interface{}) ([]byte, error) {
	return stream.Build(udsResponseBuilder, v)
}

func parseStream(data []byte, v interface{}) error {
	return stream.Parse(udsResponseBuilder, data, v)
}
