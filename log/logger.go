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
)

type Logger struct {
	Name   string
	Option Option
	Sync   bool
	logger *zap.SugaredLogger // 使用反射而不是filed字段
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
		Name:   name,
		Sync:   true,
		logger: logger.Sugar(),
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
		Name:   name,
		logger: logger.Sugar(),
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
		Name:   name,
		Sync:   true,
		logger: logger.Sugar(),
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
		Name:   name,
		logger: logger.Sugar(),
	}
}

// NewZap 不使用sugar而是使用原生的zap logger
func NewZap(c zap.Config, op ...zap.Option) (*zap.Logger, error) {
	return c.Build(op...)
}
