/*
Create: 2022/8/21
Project: FuShin
Github: https://github.com/landers1037
Copyright Renj
*/

// Package buf
package buf

import (
	"bytes"
	"encoding/gob"
	"io"
)

func decode(r io.Reader, d interface{}) error {
	return gob.NewDecoder(r).Decode(d)
}

// Decode d 需要是一个指针接收体
func Decode(b []byte, d interface{}) error {
	reader := bytes.NewReader(b)
	return decode(reader, d)
}
