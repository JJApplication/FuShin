/*
Create: 2022/8/6
Project: FuShin
Github: https://github.com/landers1037
Copyright Renj
*/

// Package log
package log

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitLogger(t *testing.T) {
	l := Logger{
		Name:   "test",
		Option: DevOption,
	}
	err := l.Init()
	if assert.NoError(t, err) {
		t.Log("test init logger with Init success")
	} else {
		t.Log("test init logger with Init failed")
	}

	// 错误的场景
	err = l.Init()
	if assert.Error(t, err) {
		if assert.Equal(t, err.Error(), "logger has already been created") {
			t.Log("test Init a logger secondly success")
		} else {
			t.Logf("test Init a logger secondly failed %v", err)
		}
	}
}

func TestLoggerOpt(t *testing.T) {
	l := Dev("test")
	// log
	l.Info("this is INFO test")
	l.InfoF("this is INFO test %s", "Hello")
	l.Warn("this is WARN test")
	l.WarnF("this is WARN test %s", "Hello")
	l.Error("this is ERROR test")
	l.ErrorF("this is ERROR test %s", "Hello")
}

func TestLog2File(t *testing.T) {
	l := Logger{
		Name: "test",
		Option: Option{
			Development:  false,
			Level:        0,
			EnableCaller: false,
			StackTrace:   false,
			Encoding:     "json",
			Output:       []string{"stderr", "test.log"},
			EncodeConfig: DefaultEncodeConfig(),
		},
		Sync: false,
	}
	l.Init()
	l.Info("Hello World")
}
