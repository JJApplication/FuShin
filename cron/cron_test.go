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

func TestScheduleJob(t *testing.T) {
	job := NewSchedule(time.Now().Add(time.Second*2), func() {
		t.Log("test run at ", time.Now().String())
	})

	job.Start()

	// 停止任务
	job.Stop()
	time.Sleep(1 * time.Second)
	// 重启任务
	job.ReStart()
	time.Sleep(time.Second * 5)
}

func TestScheduleJobTimeout(t *testing.T) {
	job := NewSchedule(time.Now().Add(time.Second*1), func() {
		t.Log("test run at ", time.Now().String())
	})

	job.Start()

	// 停止任务
	time.Sleep(100)
	job.Stop()
	time.Sleep(5 * time.Second)
	// 重启任务 此时计划任务处于过去式 不再执行
	job.ReStart()
	time.Sleep(time.Second * 10)
}
