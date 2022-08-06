/*
Create: 2022/8/6
Project: FuShin
Github: https://github.com/landers1037
Copyright Renj
*/

// Package log
package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// ConvertConfig 从fushin option转为zap config
func ConvertConfig(op Option) zap.Config {
	return zap.Config{
		Level:             zap.NewAtomicLevelAt(zapcore.Level(op.Level)),
		Development:       op.Development,
		DisableCaller:     !op.EnableCaller,
		DisableStacktrace: !op.StackTrace,
		Sampling: &zap.SamplingConfig{
			Initial:    100,
			Thereafter: 100,
		},
		Encoding:         op.Encoding,
		EncoderConfig:    zapEncodeConfig,
		OutputPaths:      op.Output,
		ErrorOutputPaths: op.Output,
	}
}
