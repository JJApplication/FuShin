/*
Create: 2022/8/11
Project: FuShin
Github: https://github.com/landers1037
Copyright Renj
*/

// Package env
package env

import (
	"strconv"
	"time"
)

// 环境变量assert

type EnvLoader struct {
	KeepRawValue bool // 断言失败时返回原始值 为false时raw永远为原始值的空值
	raw          string
}

func (e *EnvLoader) GetValue(key string) string {
	return GetEnv(key)
}

func (e *EnvLoader) clone() *EnvLoader {
	return new(EnvLoader)
}

// Get 加载环境变量
func (e *EnvLoader) Get(key string) *EnvLoader {
	clone := e.clone()
	clone.raw = GetEnv(key)
	return clone
}

// Raw 返回原始值
func (e *EnvLoader) Raw() string {
	return e.raw
}

func (e *EnvLoader) Int() int {
	res, err := strconv.Atoi(e.Raw())
	if err != nil {
		return 0
	}
	return res
}

func (e *EnvLoader) Int64() int64 {
	res, err := strconv.ParseInt(e.raw, 10, 64)
	if err != nil {
		return 0
	}
	return res
}

func (e *EnvLoader) Int32() int32 {
	res, err := strconv.ParseInt(e.raw, 10, 32)
	if err != nil {
		return 0
	}
	return int32(res)
}

func (e *EnvLoader) Int16() int16 {
	res, err := strconv.ParseInt(e.raw, 10, 16)
	if err != nil {
		return 0
	}
	return int16(res)
}

func (e *EnvLoader) Int8() int8 {
	res, err := strconv.ParseInt(e.raw, 10, 8)
	if err != nil {
		return 0
	}
	return int8(res)
}

func (e *EnvLoader) Float64() float64 {
	res, err := strconv.ParseFloat(e.raw, 64)
	if err != nil {
		return 0
	}
	return res
}

func (e *EnvLoader) Float32() float32 {
	res, err := strconv.ParseFloat(e.raw, 32)
	if err != nil {
		return 0
	}
	return float32(res)
}

func (e *EnvLoader) Bool() bool {
	res, err := strconv.ParseBool(e.raw)
	if err != nil {
		return false
	}
	return res
}

// Time 转换时间戳
func (e *EnvLoader) Time() time.Time {
	res := e.Int64()
	return time.Unix(res, 0)
}

func (e *EnvLoader) Interface() interface{} {
	return e.raw
}

// 带默认值的加载

func (e *EnvLoader) MustString(def string) string {
	if e.raw == "" {
		return def
	}
	return e.Raw()
}

func (e *EnvLoader) MustInt(def int) int {
	if e.raw == "" {
		return def
	}
	return e.Int()
}

func (e *EnvLoader) MustInt64(def int64) int64 {
	if e.raw == "" {
		return def
	}
	return e.Int64()
}

func (e *EnvLoader) MustInt32(def int32) int32 {
	if e.raw == "" {
		return def
	}
	return e.Int32()
}

func (e *EnvLoader) MustInt16(def int16) int16 {
	if e.raw == "" {
		return def
	}
	return e.Int16()
}

func (e *EnvLoader) MustInt8(def int8) int8 {
	if e.raw == "" {
		return def
	}
	return e.Int8()
}

func (e *EnvLoader) MustFloat64(def float64) float64 {
	if e.raw == "" {
		return def
	}
	return e.Float64()
}

func (e *EnvLoader) MustFloat32(def float32) float32 {
	if e.raw == "" {
		return def
	}
	return e.Float32()
}

func (e *EnvLoader) MustBool(def bool) bool {
	if e.raw == "" {
		return def
	}
	return e.Bool()
}
