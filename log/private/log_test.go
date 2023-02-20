/*
Create: 2022/8/4
Project: FuShin
Github: https://github.com/landers1037
Copyright Renj
*/

// Package internal
package private

import (
	"log"
	"testing"
)

func TestLogInfo(t *testing.T) {
	Log.Info("test1", "test2")
}

func TestLogInfoF(t *testing.T) {
	Log.InfoF("%s1", "test")
}

func TestDevLog(t *testing.T) {
	t.Log("start")
	l := NewLogger("[Test] ", log.LstdFlags, true)
	l.Error("error")
	l.Warn("warn")
	l.Info("info")
}
