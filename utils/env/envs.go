/*
Create: 2022/6/29
Project: FuShin
Github: https://github.com/landers1037
Copyright Renj
*/

// Package env
package env

import (
	"context"
	"os"
)

// GetEnv 获取环境变量
func GetEnv(key string) string {
	return os.Getenv(key)
}

// GetEnvs 获取多个环境变量
func GetEnvs(keys ...string) []string {
	var envs []string
	envs = make([]string, len(keys))
	for _, k := range keys {
		envs = append(envs, os.Getenv(k))
	}

	return envs
}

// IsEnvEqual 比较环境变量
func IsEnvEqual(env, dst string) bool {
	return os.Getenv(env) == dst
}

// SetEnvContext 设置上下文环境变量
func SetEnvContext(c context.Context, key, val string) context.Context {
	return context.WithValue(c, key, val)
}
