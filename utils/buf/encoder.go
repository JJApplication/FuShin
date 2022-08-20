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

// 编码器

func encode(w io.Writer, i interface{}) error {
	if i == nil {
		return nil
	}
	return gob.NewEncoder(w).Encode(i)
}

func Encode(d interface{}) ([]byte, error) {
	var b bytes.Buffer
	err := encode(&b, d)
	return b.Bytes(), err
}
