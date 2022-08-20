/*
Create: 2022/8/20
Project: FuShin
Github: https://github.com/landers1037
Copyright Renj
*/

// Package errors
package errors

import (
	"testing"
)

func TestFushinError(t *testing.T) {
	defer func() {
		rec := Recover(recover())
		if rec.HasRecover() {
			t.Log(rec.GetTrace())
			t.Log(rec.GID())
		} else {
			t.Log(rec)
		}
	}()
	// 模拟err
	err := New("this is an error")
	t.Log(err.Error())
	// 模拟panic
	Panic("this is a panic")
	// panic
	panic("panic2")
}

func TestFushinTryCatch(t *testing.T) {
	Try(func() {
		t.Log("start")
		panic("test panic")
	}).Catch(func(exception interface{}) {
		t.Log(exception)
	}).Finally(func() {
		t.Log("end")
	})
}
