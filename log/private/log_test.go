/*
Create: 2022/8/4
Project: FuShin
Github: https://github.com/landers1037
Copyright Renj
*/

// Package internal
package private

import (
	"testing"
)

func TestLogInfo(t *testing.T) {
	Log.info("test1", "test2")
}

func TestLogInfoF(t *testing.T) {
	Log.infoF("%s1", "test")
}
