/*
Create: 2022/8/6
Project: FuShin
Github: https://github.com/landers1037
Copyright Renj
*/

// Package uds
package uds

import (
	"testing"

	"github.com/JJApplication/fushin/server/uds"
)

func TestDial(t *testing.T) {
	client := UDSClient{Addr: "/tmp/wdnmd.sock"}
	err := client.Dial()
	if err != nil {
		t.Log(err)
		t.Skip()
	}
	err = client.Send(uds.Req{
		Operation: "testop",
		Data:      "Hello World",
		From:      "client",
		To:        nil,
	})
	if err != nil {
		t.Log(err)
		t.Skip()
	}

	t.Log("test client dial success")
	defer client.Close()
}

func TestDialWithRes(t *testing.T) {
	client := UDSClient{Addr: "/tmp/wdnmd.sock"}
	err := client.Dial()
	if err != nil {
		t.Log(err)
		t.Skip()
	}
	res, err := client.SendWithRes(uds.Req{
		Operation: "testop",
		Data:      "Hello World",
		From:      "client",
		To:        nil,
	})
	if err != nil {
		t.Log(err)
		t.Skip()
	}

	t.Logf("res is %+v\n", res)
	t.Log("test client dial with response success")
	defer client.Close()
}

func TestDialRaw(t *testing.T) {
	client := UDSClient{Addr: "/tmp/wdnmd.sock"}
	err := client.Dial()
	if err != nil {
		t.Log(err)
		t.Skip()
	}
	err = client.SendRaw("Hello World")
	if err != nil {
		t.Log(err)
		t.Skip()
	}

	t.Log("test client dial raw success")
	defer client.Close()
}
