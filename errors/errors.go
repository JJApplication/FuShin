/*
Create: 2022/8/18
Project: FuShin
Github: https://github.com/landers1037
Copyright Renj
*/

// Package errors
package errors

import (
	"github.com/google/uuid"
)

type Errors interface {
	Error() string
	ID() string
	GetTrace() string
	HasError() bool
	HasRecover() bool
}

// FushinError
// 包含goroutineID的可追溯错误
type FushinError struct {
	hasRecover bool
	message    string
	id         string
	stack      string
	gid        string
}

func New(s string) FushinError {
	return FushinError{message: s, id: uuid.New().String()}
}

func (f *FushinError) Error() string {
	return f.message
}

// ID 获取错误的uuid
func (f *FushinError) ID() string {
	return f.id
}

// GetTrace 获取stack堆栈
func (f *FushinError) GetTrace() string {
	if f.hasRecover {
		return f.stack
	}
	return ""
}

// HasError 是否有错误
func (f *FushinError) HasError() bool {
	if f.stack != "" || f.message != "" {
		return true
	}
	return false
}

func (f *FushinError) HasRecover() bool {
	if f.stack != "" && f.hasRecover {
		return true
	}
	return false
}

// GID 返回recover堆栈的goroutine id
func (f *FushinError) GID() string {
	return f.gid
}
