/*
Create: 2022/8/21
Project: FuShin
Github: https://github.com/landers1037
Copyright Renj
*/

// Package buf
package buf

import (
	"encoding/gob"
	"sync"
)

var bufMap map[interface{}]struct{}
var lock *sync.Mutex

func init() {
	lock = new(sync.Mutex)
	bufMap = make(map[interface{}]struct{})
}

// Reg 注册结构体
func Reg(d interface{}) {
	if _, ok := bufMap[d]; !ok {
		lock.Lock()
		gob.Register(d)
		bufMap[d] = struct{}{}
		lock.Unlock()
	}
}

func CheckReg(d interface{}) bool {
	if _, ok := bufMap[d]; ok {
		return true
	}
	return false
}

func List() map[interface{}]struct{} {
	return bufMap
}

type FushinBuf interface {
	Encode() ([]byte, error)
	Decode(data []byte) error
	Calc() int
}

// Buf 实现FushinBuf接口的结构体可以直接使用encode编码
type Buf struct {
}

// Encode 从b实例中编码为字节流
func (b *Buf) Encode() ([]byte, error) {
	if b == nil {
		return nil, nil
	}
	return Encode(b)
}

// Decode 从data数据中解码到b *Buf
func (b *Buf) Decode(data []byte) error {
	return Decode(data, b)
}

// Calc 计算字节流长度
func (b *Buf) Calc() int {
	if b != nil {
		data, _ := Encode(b)
		return len(data)
	}
	return 0
}
