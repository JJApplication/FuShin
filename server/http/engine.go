/*
Create: 2022/8/8
Project: FuShin
Github: https://github.com/landers1037
Copyright Renj
*/

// Package http
package http

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sort"
	"sync"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

type Server struct {
	engine       *gin.Engine       // gin engine
	mux          *sync.RWMutex     // 可重入的锁
	wrapper      []Wrapper         // 中间件 顺序加载
	router       map[string]Router // 路由
	srv          *http.Server      // 内置的http.server
	EnableLog    bool              // 使用内置的日志打印 默认输出到控制台
	Debug        bool              // 开启gin的debug
	RegSignal    []os.Signal       // 监听系统信号量
	Address      Address           // 监听地址
	Headers      map[string]string // 自定义的Headers
	Copyright    string            // 版权所有 会以header: Copyright: xx的方式返回在响应中
	MaxBodySize  int               // 最大请求体限制 默认1<<20 bytes
	ReadTimeout  int               // 继承http.Server
	WriteTimeout int               // 继承http.Server
	IdleTimeout  int               // 继承http.Server
	// todo tls
}

type Address struct {
	Host   string   // 监听的HOST
	Port   int      // 监听的端口
	Domain []string // 用于校验refer的域名(非指定域名会直接屏蔽请求) 绑定域名后会默认校验refer
}

type Router struct {
	method  string
	path    string
	wrapper []WrapperFunc
}

type RouterGroup struct {
	*gin.RouterGroup
}

type Context struct {
	*gin.Context
}

type WrapperFunc func(c *Context)

type Wrapper struct {
	Name        string
	WrapperFunc WrapperFunc
}

// Init 在其他方法被调用前初始化
func (s *Server) Init() {
	if s.Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	if s.engine == nil {
		s.engine = gin.New()
		s.engine.Use(gin.Recovery())
	}
	s.router = make(map[string]Router, 1)
	s.wrapper = make([]Wrapper, 0)
	s.mux = &sync.RWMutex{}
	s.srv = &http.Server{
		Addr:              "",
		Handler:           nil,
		TLSConfig:         nil,
		ReadTimeout:       0,
		ReadHeaderTimeout: 0,
		WriteTimeout:      0,
		IdleTimeout:       0,
		MaxHeaderBytes:    0,
	}
}

func (s *Server) Listen() error {
	if s.engine == nil {
		s.error(moduleName, ErrEngineEmpty)
		return errors.New(ErrEngineEmpty)
	} else {
		s.srv.Addr = fmt.Sprintf("%s:%d", s.Address.Host, s.Address.Port)
		s.srv.Handler = s.engine
		s.srv.ReadTimeout = time.Duration(s.ReadTimeout) * time.Second
		s.srv.WriteTimeout = time.Duration(s.WriteTimeout) * time.Second
		s.srv.IdleTimeout = time.Duration(s.IdleTimeout) * time.Second
		s.srv.MaxHeaderBytes = s.MaxBodySize
		s.RegSignals()
		s.initRegSignals()
		s.initMiddles()
		s.initRoutes()
		s.initServerConfig()
		return s.srv.ListenAndServe()
	}
}

// Run 不处理错误版本的listen
func (s *Server) Run() {
	_ = s.Listen()
}

// ListenSmooth 平滑关闭
// 在注册监听signal的时候生效 默认注册CTRL+C
func (s *Server) ListenSmooth() {
	if s.engine == nil {
		s.error(ErrEngineEmpty)
		return
	} else {
		s.srv.Addr = fmt.Sprintf("%s:%d", s.Address.Host, s.Address.Port)
		s.srv.Handler = s.engine
		s.srv.ReadTimeout = time.Duration(s.ReadTimeout) * time.Second
		s.srv.WriteTimeout = time.Duration(s.WriteTimeout) * time.Second
		s.srv.IdleTimeout = time.Duration(s.IdleTimeout) * time.Second
		s.srv.MaxHeaderBytes = s.MaxBodySize
		s.RegSignals()
		s.initRegSignals()
		s.initMiddles()
		s.initRoutes()
		s.initServerConfig()
		go func() {
			sigint := make(chan os.Signal, 1)
			signal.Notify(sigint, append([]os.Signal{}, s.RegSignal...)...)
			sig := <-sigint
			// 避免换行
			if s.EnableLog {
				fmt.Println()
			}
			s.infoF("%s signal %s received, server is closed\n", moduleName, sig.String())
			if err := s.srv.Shutdown(context.Background()); err != nil {
				s.errorF("%s server shutdown error: %s\n", moduleName, err.Error())
				return
			}
		}()

		if err := s.srv.ListenAndServe(); err != nil {
			if err == http.ErrServerClosed {
				s.info(moduleName, ErrServerClosed)
				return
			}
			s.error(moduleName, err.Error())
		}
	}
}

// RegSignals 信号注册
func (s *Server) RegSignals(sigs ...os.Signal) {
	// 自动去重
	currentSigs := s.RegSignal
	allSigs := append(currentSigs, sigs...)
	temp := make(map[string]os.Signal, 1)
	// 注册默认的interrupt信号
	temp[syscall.SIGINT.String()] = syscall.SIGINT
	for _, sig := range allSigs {
		if _, ok := temp[sig.String()]; !ok {
			temp[sig.String()] = sig
		}
	}
	// 去重后对s.regsignal重新赋值
	var newSigs []os.Signal
	for _, sig := range temp {
		newSigs = append(newSigs, sig)
	}
	s.mux.RLock()
	s.RegSignal = newSigs
	s.mux.RUnlock()
}

// 初始化信号量
func (s *Server) initRegSignals() {
	for _, sig := range s.RegSignal {
		s.infoF("%s signal [%s] registered\n", moduleName, sig.String())
	}
}

// 初始化服务器配置
func (s *Server) initServerConfig() {
	s.info(moduleName, "server init success")
	s.infoF("%s server will listen on %s:%d", moduleName, s.Address.Host, s.Address.Port)
	s.infoF("%s server domain name is %v", moduleName, s.Address.Domain)
	s.infoF("%s server debug: %v", moduleName, s.Debug)
	s.infoF("%s server enableLog: %v", moduleName, s.EnableLog)
}

// 初始化注册的路由
func (s *Server) initRoutes() {
	var routers []Router
	for _, r := range s.router {
		routers = append(routers, r)
	}
	sort.SliceStable(routers, func(i, j int) bool {
		return routers[i].path < routers[j].path
	})
	for _, r := range routers {
		s.infoF("%s [%-4s] %-20s --> route wraps with (%d wrappers)", moduleName, r.method, r.path, len(r.wrapper))
		s.engine.Handle(r.method, r.path, convertWraps(r.wrapper...)...)
	}
}

// 初始化全部的中间件
// 全局的中间件在路由初始化前调用
func (s *Server) initMiddles() {
	for _, m := range s.wrapper {
		s.infoF("%s middleware [%s] registered", moduleName, m.Name)
		if m.WrapperFunc != nil {
			s.engine.Use(convertWrap(m.WrapperFunc))
		}
	}
}

// Group 路由分组
func (s *Server) Group(path string, wrap ...WrapperFunc) *RouterGroup {
	return &RouterGroup{s.engine.Group(path, convertWraps(wrap...)...)}
}

// 路由方法 不提供语法糖写法
// 必须指定请求方法

func (s *Server) Route(method, uri string, wrap ...WrapperFunc) {
	s.mux.RLock()
	if _, ok := s.router[method+uri]; ok {
		s.errorF("%s route [%s] [%s] has already been registered\n", moduleName, method, uri)
		return
	}
	s.router[method+uri] = Router{
		method:  method,
		path:    uri,
		wrapper: wrap,
	}
	defer s.mux.RUnlock()
}

// RegMiddle 注册中间件
func (s *Server) RegMiddle(mds ...Wrapper) {
	s.mux.RLock()
	// 去重
	var m []Wrapper
	if len(s.wrapper) == 0 {
		m = mds
	} else {
		for i, _ := range mds {
			v := false
			for j, _ := range s.wrapper {
				if s.wrapper[j].Name == mds[i].Name {
					break
				}
				v = true
			}
			if v {
				m = append(m, mds[i])
			}
		}
	}

	s.wrapper = append(s.wrapper, m...)
	s.mux.RUnlock()
}
