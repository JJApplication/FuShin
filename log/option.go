/*
Create: 2022/8/6
Project: FuShin
Github: https://github.com/landers1037
Copyright Renj
*/

// Package log
package log

import (
	"go.uber.org/zap/zapcore"
)

// Option 日志的配置项
type Option struct {
	Development  bool
	Level        Level
	EnableCaller bool
	StackTrace   bool
	Encoding     string
	Output       []string
	EncodeConfig zapcore.EncoderConfig // 原生的zap配置
}

// DefaultOption 默认配置
var DefaultOption = Option{
	Development:  false,
	Level:        InfoLevel,
	StackTrace:   false,
	Encoding:     EncodingJSON,
	Output:       DefaultOutput,
	EncodeConfig: zapEncodeConfig,
}

var DevOption = Option{
	Development:  true,
	Level:        DebugLevel,
	StackTrace:   true,
	Encoding:     EncodingJSON,
	Output:       DefaultOutput,
	EncodeConfig: zapEncodeConfig,
}

var zapEncodeConfig = zapcore.EncoderConfig{
	EncodeName:       nil,
	ConsoleSeparator: " ",
	TimeKey:          "time",
	LevelKey:         "level",
	NameKey:          "name",
	CallerKey:        "caller",
	MessageKey:       "message",
	StacktraceKey:    "stack",
	LineEnding:       zapcore.DefaultLineEnding,
	EncodeLevel:      zapcore.CapitalLevelEncoder,
	EncodeTime:       zapcore.TimeEncoderOfLayout("2006-1-2 15:04:05"),
	EncodeDuration:   zapcore.StringDurationEncoder,
	EncodeCaller:     zapcore.ShortCallerEncoder,
}

func DefaultEncodeConfig() zapcore.EncoderConfig {
	return zapEncodeConfig
}

// 输出格式

const EncodingJSON = "json"
const EncodingConsole = "console"

var DefaultOutput = []string{"stderr"}
