/*
Create: 2022/8/5
Project: FuShin
Github: https://github.com/landers1037
Copyright Renj
*/

// Package uds
package uds

import (
	"net"
)

// 返回数据的格式 默认都是rest的json

type Res struct {
	Error string   `json:"error"`
	Data  string   `json:"data"`
	From  string   `json:"from,omitempty"`
	To    []string `json:"to,omitempty"` // 可以有多个接收者
}

type Req struct {
	Operation string   `json:"operation"`
	Data      string   `json:"data"`
	From      string   `json:"from,omitempty"`
	To        []string `json:"to,omitempty"` // 可以有多个接收者
}

type Funcs map[string]func(c net.Conn, req Req)
