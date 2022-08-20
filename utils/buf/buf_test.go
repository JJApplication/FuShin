/*
Create: 2022/8/21
Project: FuShin
Github: https://github.com/landers1037
Copyright Renj
*/

// Package buf
package buf

import (
	"testing"
)

type TestStruct struct {
	Name string
}

func TestBuf(t *testing.T) {
	Reg(TestStruct{})
	t.Logf("%+v", List())
	data := TestStruct{
		Name: "yes",
	}

	res, err := Encode(data)
	t.Log(res, err)

	var tmp TestStruct
	err = Decode(res, &tmp)
	t.Log(tmp, err)
}

type TestBufExtend struct {
	Buf
	Name string
}

func TestBufFunc(t *testing.T) {
	tb := new(TestBufExtend)
	tb.Name = "ok"
	res, err := tb.Encode()
	t.Log(res, err)

	err = tb.Decode(res)
	t.Log(tb, err)

	t.Log(tb.Calc())
}
