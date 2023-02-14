/*
Create: 2022/8/5
Project: FuShin
Github: https://github.com/landers1037
Copyright Renj
*/

// Package log
package log

import (
	"errors"
	"fmt"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	Name      string
	Option    Option
	Sync      bool
	logger    *zap.SugaredLogger // 使用反射而不是filed字段
	zapLogger *zap.Logger
}

// Init 从Logger内置的参数新建日志构建logger
func (l *Logger) Init() error {
	if l.logger != nil {
		return errors.New("logger has already been created")
	}
	lock := sync.Mutex{}
	lock.Lock()
	logger, err := ConvertConfig(l.Option).Build(zap.AddCaller(), zap.AddCallerSkip(1))
	if err != nil {
		fmt.Printf("%s create error: %s\n", moduleName, err)
		return err
	}
	if l.Sync {
		defer logger.Sync()
	}
	l.logger = logger.Sugar()
	defer lock.Unlock()
	return nil
}

func (l Logger) Info(v ...interface{}) {
	l.logger.Info(v...)
}

func (l Logger) InfoF(fmtStr string, v ...interface{}) {
	l.logger.Infof(fmtStr, v...)
}

func (l Logger) Warn(v ...interface{}) {
	l.logger.Warn(v...)
}

func (l Logger) WarnF(fmtStr string, v ...interface{}) {
	l.logger.Warnf(fmtStr, v...)
}

func (l Logger) Error(v ...interface{}) {
	l.logger.Error(v...)
}

func (l Logger) ErrorF(fmtStr string, v ...interface{}) {
	l.logger.Errorf(fmtStr, v...)
}

func Default(name string) *Logger {
	logger, err := ConvertConfig(DefaultOption).Build(zap.AddCaller(), zap.AddCallerSkip(1))
	if err != nil {
		fmt.Printf("%s create error: %s\n", moduleName, err)
		return &Logger{}
	}
	defer logger.Sync()
	return &Logger{
		Name:      name,
		Sync:      true,
		logger:    logger.Sugar(),
		zapLogger: logger,
	}
}

func Dev(name string) *Logger {
	logger, err := ConvertConfig(DevOption).Build(zap.AddCaller(), zap.AddCallerSkip(1))
	if err != nil {
		fmt.Printf("%s create error: %s\n", moduleName, err)
		return &Logger{}
	}
	defer logger.Sync()
	return &Logger{
		Name:      name,
		logger:    logger.Sugar(),
		zapLogger: logger,
	}
}

// New 新建日志
func New(name string, op Option) *Logger {
	logger, err := ConvertConfig(op).Build(zap.AddCaller(), zap.AddCallerSkip(1))
	if err != nil {
		fmt.Printf("%s create error: %s\n", moduleName, err)
		return &Logger{}
	}
	defer logger.Sync()
	return &Logger{
		Name:      name,
		Sync:      true,
		logger:    logger.Sugar(),
		zapLogger: logger,
	}
}

// Raw 使用zap原生配置创建
func Raw(name string, c zap.Config, op ...zap.Option) Logger {
	logger, err := c.Build(op...)
	if err != nil {
		fmt.Printf("%s create error: %s\n", moduleName, err)
		return Logger{}
	}
	return Logger{
		Name:      name,
		logger:    logger.Sugar(),
		zapLogger: logger,
	}
}

// NewZap 不使用sugar而是使用原生的zap logger
func NewZap(c zap.Config, op ...zap.Option) (*zap.Logger, error) {
	return c.Build(op...)
}

// NewCore 基于core创建
func NewCore(name string, op Option) *Logger {
	zapCore := zapcore.NewCore(
		coreEncoder(op.Encoding),
		zapcore.AddSync(wrapRotateCore(name, op.RotateOption)),
		zap.NewAtomicLevel(),
	)

	logger := zap.New(zapCore, zap.AddCaller(), zap.AddCallerSkip(1))
	return &Logger{
		Name:      name,
		Sync:      true,
		logger:    logger.Sugar(),
		zapLogger: logger,
	}
}
