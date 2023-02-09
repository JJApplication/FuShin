/*
Create: 2023/2/9
Project: FuShin
Github: https://github.com/landers1037
Copyright Renj

基于endless提供go server的平滑启动
*/

package endless

import (
	"context"
	"net/http"
	"time"

	fushinHttp "github.com/JJApplication/fushin/server/http"
)

type EndlessConfig struct {
	Addr         string
	Handler      http.Handler
	IdleTimeout  time.Duration
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	BeforeBegin  func(add string)
}

type EndlessServer struct {
	srv *endlessServer
}

// EndlessRun 直接对endless的快速调用
func EndlessRun(address string, hd http.Handler) error {
	return ListenAndServe(address, hd)
}

// EndlessFushin 直接对endless的快速调用 以fushinServer初始化
func EndlessFushin(srv *fushinHttp.Server) error {
	return ListenAndServe(srv.GetAddr(), srv.GetEngine())
}

// EndlessHooks 直接对endless的快速调用 增加必备hooks
func EndlessHooks(addr string, hd http.Handler, beforeHooks func(add string)) error {
	s := NewServer(addr, hd)
	s.BeforeBegin = beforeHooks
	return s.ListenAndServe()
}

// NewEndless 新建endless服务
func NewEndless(cf EndlessConfig) *EndlessServer {
	s := NewServer(cf.Addr, cf.Handler)
	s.BeforeBegin = cf.BeforeBegin
	if cf.IdleTimeout > 0 {
		s.IdleTimeout = cf.IdleTimeout
	}

	if cf.ReadTimeout > 0 {
		s.ReadTimeout = cf.ReadTimeout
	}

	if cf.WriteTimeout > 0 {
		s.WriteTimeout = cf.WriteTimeout
	}

	s.BeforeBegin = cf.BeforeBegin

	return &EndlessServer{srv: s}
}

// Shutdown for endlessServer
func (e *EndlessServer) Shutdown() {
	e.srv.shutdown()
}

// ShutdownCtx for endlessServer
func (e *EndlessServer) ShutdownCtx(ctx context.Context) error {
	return e.srv.Shutdown(ctx)
}

// RegHooksShutdown Register hooks on Shutdown for endlessServer
func (e *EndlessServer) RegHooksShutdown(f func()) {
	e.srv.RegisterOnShutdown(f)
}

// Close for endlessServer
func (e *EndlessServer) Close() error {
	return e.srv.Close()
}
