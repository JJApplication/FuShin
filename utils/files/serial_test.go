/*
Create: 2022/8/18
Project: FuShin
Github: https://github.com/landers1037
Copyright Renj
*/

// Package files
package files

import (
	"testing"
	"time"
)

func TestSerial(t *testing.T) {
	s := NewSerial("")
	s.WriteSerial("Test start\n")
	s.WriteSerial(time.Now().Format("2006-01-02-15:04:05"))
	s.WriteSerial("\nTest end\n")

	_, err := s.ExportSerial()
	if err != nil {
		t.Error(err)
	}
	err = s.ToFileSerial()
	t.Log(err)
}
