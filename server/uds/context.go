/*
Create: 2022/8/12
Project: FuShin
Github: https://github.com/landers1037
Copyright Renj
*/

// Package uds
package uds

import (
	"net"
)

// UDSContext UDS服务端的上下文
type UDSContext struct {
	operation string
	c         net.Conn
}

// Operation 获取注册到此上下文的操作名称
func (uc *UDSContext) Operation() string {
	return uc.operation
}

// Response 响应Res
func (uc *UDSContext) Response(res Res) error {
	return Response(uc.c, res)
}

// ResponseAny 响应任意数据
func (uc *UDSContext) ResponseAny(res interface{}) error {
	return ResponseAny(uc.c, res)
}
