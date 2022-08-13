/*
Create: 2022/8/5
Project: FuShin
Github: https://github.com/landers1037
Copyright Renj
*/

// Package uds
package uds

import (
	"io"
	"net"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/JJApplication/fushin/log"
	"github.com/JJApplication/fushin/utils/json"
)

// 连接

type UDSServer struct {
	Name      string              // 注册服务的名称
	Option    Option              // 默认使用fushin option
	Logger    log.LoggerInterface // 默认使用fushin logger
	listener  *net.UnixListener   // 内部的listener
	closeFlag chan int
	funcs     Funcs // 注册的操作
}

// New 返回全新的unix server
func New(name string) *UDSServer {
	return &UDSServer{
		Name:      name,
		Option:    Option{MaxSize: 1 << 10},
		Logger:    nil,
		closeFlag: make(chan int),
	}
}

// NewWithOption 返回自定义的unix server
func NewWithOption(name string, o Option, l log.LoggerInterface) *UDSServer {
	return &UDSServer{
		Name:      name,
		Option:    o,
		Logger:    l,
		closeFlag: make(chan int),
	}
}

// Default 返回默认的unix server
func Default(name string) *UDSServer {
	return &UDSServer{
		Name:      name,
		Option:    DefaultOption,
		Logger:    nil,
		closeFlag: make(chan int),
	}
}

// Listen 在协程启动监听
func (u *UDSServer) Listen() error {
	if u.Option.MaxSize <= 0 {
		u.error(moduleName, "uds server request maxsize not set")
	}
	addr, err := net.ResolveUnixAddr("unix", GetSocket(u.Name))
	if err != nil {
		u.errorF("%s resolve unix address error: %s", moduleName, err.Error())
		return err
	}
	listener, err := net.ListenUnix("unix", addr)
	if err != nil {
		u.errorF("%s listen on unix address [%s] error: %s", moduleName, addr, err.Error())
		return err
	}

	u.listener = listener
	go u.AutoCheck()
	go u.runtimeClosed()
	for {
		conn, err := u.listener.Accept()
		if err != nil && strings.Contains(err.Error(), net.ErrClosed.Error()) {
			break
		}
		if err != nil {
			u.errorF("%s unix accept error: %s", moduleName, err.Error())
			continue
		}
		go u.Run(conn)
	}

	return err
}

// Run 启动服务监听请求
func (u *UDSServer) Run(c net.Conn) {
	if u.Option.AutoRecover {
		defer func() {
			if err := recover(); err != nil {
				u.errorF("%s [recover] from uds server: %v", moduleName, err)
			}
		}()
	}

	u.infoF("%s uds client [%s] connected", moduleName, c.RemoteAddr().String())

	for {
		buf := make([]byte, u.Option.MaxSize)
		count, err := c.Read(buf)
		if err != nil {
			if err == io.EOF {
				u.errorF("%s read from uds client, client disconnected", moduleName)
				break
			}
			u.errorF("%s read from uds client error: %s", moduleName, err.Error())
			continue
		}
		if count <= 1 {
			u.error(moduleName, "received message is null")
			continue
		}

		// 开发模式下打印报文
		u.infoF("%s message [%s] received", moduleName, strings.TrimSuffix(string(buf[:count]), ""))
		reqBody := u.Option.RequestFormat
		err = json.Json.Unmarshal(buf[:count], &reqBody)
		if err != nil {
			u.errorF("%s decode request from client error: %s", moduleName, err.Error())
			continue
		}
		// 匹配操作
		if f, ok := u.funcs[reqBody.Operation]; ok {
			f(&UDSContext{c: c, operation: reqBody.Operation}, reqBody)
		} else {
			u.warnF("%s unsupported operation [%s]", moduleName, reqBody.Operation)
			u.warnF("%s server response with error: %v", moduleName, Response(c, Res{
				Error: ErrUnsupportedOperation,
				Data:  "",
				From:  "",
				To:    nil,
			}))
			continue
		}
	}
	defer c.Close()
}

// AddFunc 增加uds响应的处理函数
// 响应会以注册的format格式化
// 针对不同的处理可以注册多个func处理请求
// operation为请求中的操作关键字 用于匹配func
func (u *UDSServer) AddFunc(operation string, f func(c *UDSContext, req Req)) {
	// insert operation name
	if u.funcs == nil {
		u.funcs = make(Funcs, 1)
	}
	if operation != "" {
		u.funcs[operation] = f

	} else {
		u.error(moduleName, "addFunc operation is empty")
	}
}

// Close 关闭unix的listener 不再接收请求
func (u *UDSServer) Close() {
	lock := sync.Mutex{}
	lock.Lock()
	u.listener.Close()
	u.closeFlag <- 1
	lock.Unlock()
}

func (u *UDSServer) runtimeClosed() {
	for {
		select {
		case <-u.closeFlag:
			u.info(moduleName, "unix server close signal received")
			u.info(moduleName, "unix server is closed")
			close(u.closeFlag)
			os.Exit(0)
		default:
			if u.listener == nil {
				u.info(moduleName, "unix server is closed")
				close(u.closeFlag)
				os.Exit(1)
			}
		}
	}
}

func (u *UDSServer) AutoCheck() {
	c := time.Tick(time.Duration(u.Option.AutoCheckDuration) * time.Second)
	for range c {
		if u.listener == nil {
			u.error(moduleName, "autoCheck unix listener is nil")
		} else {
			u.info(moduleName, "autoCheck unix listener is good")
		}
	}
}
