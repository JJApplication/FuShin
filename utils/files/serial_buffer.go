/*
Create: 2022/8/18
Project: FuShin
Github: https://github.com/landers1037
Copyright Renj
*/

// Package files
package files

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"time"
)

// 串行的流写入
// 适用于同步执行的调用，以串行的方式记录调用信息
// 最终记录完毕后 可以统一输出流数据

type Serial struct {
	name string
	buf  bytes.Buffer
}

// NewSerial 新建一个串行的流
// 无name时 最终会默认增加一个时间戳
func NewSerial(name string) *Serial {
	buf := bytes.Buffer{}
	return &Serial{name: name, buf: buf}
}

// WriteSerial 写入可读性强的string
func (s *Serial) WriteSerial(str string) {
	s.buf.WriteString(str)
}

// WriteSerialRaw 写入bytes
func (s *Serial) WriteSerialRaw(b []byte) {
	s.buf.Write(b)
}

// GetSerial 获取内容
func (s *Serial) GetSerial() string {
	return s.buf.String()
}

// ExportSerial 导出流到指定writer
// 默认为os.stdout
func (s *Serial) ExportSerial(w ...io.Writer) (int, error) {
	if len(w) == 0 {
		return fmt.Fprint(os.Stdout, s.buf.String())
	} else {
		return fmt.Fprint(io.MultiWriter(w...), s.buf.String())
	}
}

// ToFileSerial 存储到文件
// 默认存储为s.name.txt || timestamp.txt
func (s *Serial) ToFileSerial() error {
	content := s.GetSerial()
	if s.name == "" {
		s.name = time.Now().Format("2006-01-02-15-04-05.00000")
	}
	return ioutil.WriteFile(s.name+".txt", []byte(content), 0644)
}
