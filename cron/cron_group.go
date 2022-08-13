/*
Create: 2022/8/14
Project: FuShin
Github: https://github.com/landers1037
Copyright Renj
*/

// Package cron
package cron

import (
	"sync"

	"github.com/robfig/cron/v3"
)

// CronGroup 相同定时的任务组
type CronGroup struct {
	Spec  string // cronjob规则
	c     *cron.Cron
	lock  *sync.Mutex
	count int
}

// NewGroup 新建一个规则分组
func NewGroup(spec string) *CronGroup {
	return &CronGroup{
		Spec:  spec,
		lock:  new(sync.Mutex),
		c:     cron.New(cron.WithSeconds()),
		count: 0,
	}
}

// AddFunc 添加一个定时任务
func (g *CronGroup) AddFunc(f func()) (int, error) {
	id, err := g.c.AddFunc(g.Spec, f)
	if err == nil {
		g.addCount()
	}
	return int(id), err
}

func (g *CronGroup) Start() {
	g.c.Start()
}

func (g *CronGroup) Stop() {
	g.c.Stop()
}

func (g *CronGroup) Cancel(id int) {
	g.subCount()
	g.c.Remove(cron.EntryID(id))
}

func (g *CronGroup) addCount() {
	g.lock.Lock()
	g.count += 1
	g.lock.Unlock()
}

func (g *CronGroup) subCount() {
	g.lock.Lock()
	g.count -= 1
	g.lock.Unlock()
}

func (g *CronGroup) JobsCount() int {
	return g.count
}

func (g *CronGroup) JobSpec() string {
	return g.Spec
}
