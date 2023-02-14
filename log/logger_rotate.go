/*
Create: 2023/2/10
Project: FuShin
Github: https://github.com/landers1037
Copyright Renj
*/

package log

import (
	"fmt"

	"github.com/gookit/goutil/reflects"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// 增加lumberkack的日志分割

func wrapRotateCore(name string, op RotateOption) *lumberjack.Logger {
	if reflects.IsEqual(op, RotateOption{}) {
		op = DefaultRotateOption
	}
	return &lumberjack.Logger{
		Filename:   fmt.Sprintf("%s.log", name),
		MaxSize:    op.MaxSize,
		MaxAge:     op.MaxAge,
		MaxBackups: op.MaxBackups,
		LocalTime:  op.LocalTime,
		Compress:   op.Compress,
	}
}

func coreConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		TimeKey:             "time",
		LevelKey:            "level",
		NameKey:             "name",
		CallerKey:           "caller",
		MessageKey:          "message",
		StacktraceKey:       "stack",
		SkipLineEnding:      false,
		LineEnding:          zapcore.DefaultLineEnding,
		EncodeLevel:         zapcore.CapitalLevelEncoder,
		EncodeTime:          zapcore.TimeEncoderOfLayout("2006-1-2 15:04:05"),
		EncodeDuration:      zapcore.StringDurationEncoder,
		EncodeCaller:        zapcore.ShortCallerEncoder,
		EncodeName:          nil,
		NewReflectedEncoder: nil,
		ConsoleSeparator:    "",
	}
}

func coreEncoder(typ string) zapcore.Encoder {
	switch typ {
	case EncodingJSON:
		return zapcore.NewJSONEncoder(coreConfig())
	case EncodingConsole:
		return zapcore.NewConsoleEncoder(coreConfig())
	default:
		return zapcore.NewConsoleEncoder(coreConfig())
	}
}
