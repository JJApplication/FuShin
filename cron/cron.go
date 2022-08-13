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

type Cron struct {
	c    *cron.Cron
	g    map[int]string // 内部维护的任务组
	lock *sync.Mutex
}

// New 创建支持秒级的定时器
func New() *Cron {
	return &Cron{
		c:    cron.New(cron.WithSeconds()),
		g:    make(map[int]string, 1),
		lock: new(sync.Mutex),
	}
}

// NewWithRecover 创建带recover的支持秒级的定时器
func NewWithRecover() *Cron {
	return &Cron{
		c:    cron.New(cron.WithSeconds(), cron.WithChain(cron.Recover(cron.DefaultLogger))),
		g:    make(map[int]string, 1),
		lock: new(sync.Mutex),
	}
}

// NewWithOption 附加原生cron v3配置
func NewWithOption(ops ...cron.Option) *Cron {
	return &Cron{
		c:    cron.New(ops...),
		g:    make(map[int]string, 1),
		lock: new(sync.Mutex),
	}
}

// AddFunc 添加规则
// spec 为标准的cronjob表达式
// 返回为唯一的定时id和错误
func (cronJob *Cron) AddFunc(spec string, f func()) (int, error) {
	id, err := cronJob.c.AddFunc(spec, f)
	if err == nil {
		cronJob.lock.Lock()
		cronJob.g[int(id)] = spec
		cronJob.lock.Unlock()
	}
	return int(id), err
}

// Start 启动cronJob
func (cronJob *Cron) Start() {
	cronJob.c.Start()
}

// Stop 停止cronJob
func (cronJob *Cron) Stop() {
	cronJob.c.Stop()
}

// Cancel 取消指定的任务
func (cronJob *Cron) Cancel(id int) {
	if _, ok := cronJob.g[id]; ok {
		cronJob.lock.Lock()
		delete(cronJob.g, id)
		cronJob.lock.Unlock()
	}
	cronJob.c.Remove(cron.EntryID(id))
}

// Jobs 展示当前加入的任务
func (cronJob *Cron) Jobs() map[int]string {
	return cronJob.g
}

// JobSpec 返回指定任务的规则
func (cronJob *Cron) JobSpec(id int) string {
	if _, ok := cronJob.g[id]; ok {
		return cronJob.g[id]
	}
	return ""
}
