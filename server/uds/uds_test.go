/*
Create: 2022/8/5
Project: FuShin
Github: https://github.com/landers1037
Copyright Renj
*/

// Package uds
package uds

import (
	"context"
	"testing"
	"time"
)

// 服务端测试

func TestNewUDSServerExit(t *testing.T) {
	server := Default("/tmp/Test")

	go server.Listen()
	defer server.Close()
	time.Sleep(2 * time.Second)
}

func TestNewUDSServer(t *testing.T) {
	server := Default("/tmp/Test")
	server.AddFunc("op1", func(c *UDSContext, req Req) {
		t.Logf("rec %+v", req)
		err := Response(c.c, Res{
			Error: "",
			Data:  "Hello World",
			From:  "",
			To:    nil,
		})
		t.Log(err)
	})

	go server.Listen()
	c, cancel := context.WithTimeout(context.Background(), time.Second*60)
	defer cancel()
	for {
		select {
		case <-c.Done():
			t.Log("exit server")
			server.Close()
		default:
		}
	}
}
