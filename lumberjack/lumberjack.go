/*
   Create: 2023/8/19
   Project: FuShin
   Github: https://github.com/landers1037
   Copyright Renj
*/

// Package lumberjack
//
// log rotate interface
package lumberjack

import "gopkg.in/natefinch/lumberjack.v2"

type Option struct {
	Filename   string
	MaxSize    int
	MaxAge     int
	MaxBackups int
	LocalTime  bool
	Compress   bool
}

func NewEmptyRotate() Rotate {

	return &lumberjack.Logger{}
}

func NewRotate(option Option) Rotate {
	return &lumberjack.Logger{
		Filename:   option.Filename,
		MaxSize:    option.MaxSize,
		MaxAge:     option.MaxAge,
		MaxBackups: option.MaxBackups,
		LocalTime:  option.LocalTime,
		Compress:   option.Compress,
	}
}
