/*
Create: 2022/8/14
Project: FuShin
Github: https://github.com/landers1037
Copyright Renj
*/

// Package cron
package cron

import (
	"math"
	"time"
)

// 定时计划任务
// 仅在特定的时间执行
// 时间使用time.Time

type ScheduleJob struct {
	ExecTime time.Time    // 执行时间
	Name     string       // 可选 任务名称
	tick     *time.Ticker // 默认的定时1s 太多定时器会导致协程压力过大
	done     bool
	stop     bool
	start    bool
	f        func()
}

// NewSchedule 创建一个计划任务
// 给定的时间如果在过去或等于现在的时间 不做任何事情
func NewSchedule(t time.Time, f func()) *ScheduleJob {
	return &ScheduleJob{
		ExecTime: t,
		Name:     "",
		tick:     time.NewTicker(time.Second),
		done:     false,
		stop:     false,
		f:        f,
	}
}

func (s *ScheduleJob) SetName(name string) {
	s.Name = name
}

func (s *ScheduleJob) Start() {
	if !s.start {
		s.start = true
		go func() {
			for {
				select {
				case <-s.tick.C:
					if calc(s.ExecTime) {
						s.f()
						// done
						s.Stop()
					}
				}
			}
		}()
	}
}

func (s *ScheduleJob) Stop() {
	s.tick.Stop()
	s.stop = true
}

// ReStart 从停止状态下重新启动
func (s *ScheduleJob) ReStart() {
	if !s.done && s.stop {
		s.tick.Reset(time.Second)
		s.stop = false
	}
}

// IsStart 任务是否启动
func (s *ScheduleJob) IsStart() bool {
	return s.start
}

// IsDone 是否执行完毕
func (s *ScheduleJob) IsDone() bool {
	return s.done
}

// IsStopped 是否停止
func (s *ScheduleJob) IsStopped() bool {
	return s.stop
}

func now() time.Time {
	return time.Now()
}

// 计算时间合法性 对于秒级任务 可能定时器刷新不及时 +-2s都算执行时间
// 非法时返回false
func calc(t time.Time) bool {
	n := now()
	if math.Abs(n.Sub(t).Seconds()) <= 2 {
		return true
	}
	return false
}
