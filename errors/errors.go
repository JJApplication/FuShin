/*
Create: 2022/8/18
Project: FuShin
Github: https://github.com/landers1037
Copyright Renj
*/

// Package errors
package errors

type Errors interface {
	Error() string
	ID() string
	GetTrace()
}
