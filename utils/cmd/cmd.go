/*
Create: 2022/8/24
Project: FuShin
Github: https://github.com/landers1037
Copyright Renj

cmd 命令行exec工具类
*/

// Package cmd
package cmd

import (
	"context"
	"os/exec"
)

// 原生的执行命令 直接使用cmd.exec实现
// 此部分只封装常用的shell调用

const (
	Bash = "bash"
	Zsh  = "zsh"
	Fish = "fish"
	Sh   = "sh"
)

const (
	DefaultMask = 0117 // 默认0660
)

type Cmder struct {
	shellUse string // 默认使用的shell
}

var Cmd = Cmder{shellUse: Bash}

func InitCmder(shell string) {
	Cmd.shellUse = shell
}

// NewCmder default BASH
func NewCmder() Cmder {
	return Cmder{shellUse: Bash}
}

func New(shell string, mask int) Cmder {
	if shell == "" {
		return Cmder{Bash}
	}
	return Cmder{shellUse: shell}
}

func (c Cmder) fmt(sh string) []string {
	if !check(sh) {
		return nil
	}
	return []string{"-c", sh}
}

func (c Cmder) Exec(sh string) error {
	cmd := exec.Command(c.shellUse, c.fmt(sh)...)
	return cmd.Run()
}

func (c Cmder) ExecOpt(sh string) ([]byte, error) {
	cmd := exec.Command(c.shellUse, c.fmt(sh)...)
	return cmd.Output()
}

func (c Cmder) ExecStr(sh string) (string, error) {
	cmd := exec.Command(c.shellUse, c.fmt(sh)...)
	d, e := cmd.Output()
	return string(d), e
}

func (c Cmder) ExecCtx(sh string, ctx context.Context) ([]byte, error) {
	cmd := exec.CommandContext(ctx, c.shellUse, c.fmt(sh)...)
	return cmd.Output()
}

func (c Cmder) ExecAsync(sh string, ctx context.Context) error {
	cmd := exec.CommandContext(ctx, c.shellUse, c.fmt(sh)...)
	return cmd.Start()
}

// Run 返回原生的cmd
func (c Cmder) Run(sh string) *exec.Cmd {
	return exec.Command(c.shellUse, c.fmt(sh)...)
}

func (c Cmder) RunContext(sh string, ctx context.Context) *exec.Cmd {
	return exec.CommandContext(ctx, c.shellUse, c.fmt(sh)...)
}
