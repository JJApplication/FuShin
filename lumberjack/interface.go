/*
   Create: 2023/8/19
   Project: FuShin
   Github: https://github.com/landers1037
   Copyright Renj
*/

package lumberjack

// interface for rotate

type Rotate interface {
	Write(p []byte) (n int, err error)
	Close() error
	Rotate() error
}
