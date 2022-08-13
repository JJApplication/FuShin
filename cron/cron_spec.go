/*
Create: 2022/8/14
Project: FuShin
Github: https://github.com/landers1037
Copyright Renj
*/

// Package cron
package cron

import (
	"fmt"
)

// 常用的cron规则

const (
	Yearly     = "@yearly"
	Annually   = "@annually"
	Monthly    = "@monthly"
	Daily      = "@daily"
	MidNightly = "@midnight"
	Hourly     = "@hourly"
	Weekly     = "@weekly"
	Every      = "@every" // 后面使用秒s 分m 时h
)

// EveryFmt 每分或每秒
// 例如1s
// 例如5m
// 例如 1h1m
func EveryFmt(t string) string {
	return fmt.Sprintf("%s %s", Every, t)
}
