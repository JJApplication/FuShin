/*
Create: 2022/8/14
Project: FuShin
Github: https://github.com/landers1037
Copyright Renj
*/

// Package cron
package cron

import (
	"testing"
	"time"
)

func TestCron(t *testing.T) {
	c := New()
	c.AddFunc("@every 1s", func() {
		t.Log("test func run")
	})
	c.Start()

	time.Sleep(time.Second * 10)
	c.Stop()
}

func TestCronGroup(t *testing.T) {
	cg := NewGroup("@every 1s")
	cg.AddFunc(func() {
		t.Log("test func run")
	})
	cg.Start()

	time.Sleep(time.Second * 10)
	cg.Stop()
}

func TestCronSpec(t *testing.T) {
	cg := NewGroup(EveryFmt("5s"))
	cg.AddFunc(func() {
		t.Log("test func run")
	})
	cg.Start()

	time.Sleep(time.Second * 10)
	cg.Stop()
}
