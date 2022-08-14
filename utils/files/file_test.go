/*
Create: 2022/8/14
Project: FuShin
Github: https://github.com/landers1037
Copyright Renj
*/

// Package files
package files

import (
	"testing"
)

func TestFileInfo(t *testing.T) {
	testFile := "file_stat.go"
	t.Log(IsFile(testFile))
	t.Log(IsFolder(testFile))
	t.Log(IsExist(testFile))
	t.Log(GetModTime(testFile))
	t.Log(GetSize(testFile))
	t.Log(GetSizeReadable(testFile))
	t.Log(GetFileMode(testFile))
}

func TestZip(t *testing.T) {
	testFile := "file_stat.go"
	err := Zip(testFile, "file.zip")
	t.Log(err)
}
