/*
Create: 2022/8/9
Project: FuShin
Github: https://github.com/landers1037
Copyright Renj
*/

// Package http
package http

import (
	"fmt"
	"os"
	"syscall"
	"testing"
)

func TestNewServer(t *testing.T) {
	s := Server{
		Debug:        true,
		EnableLog:    false,
		RegSignal:    nil,
		Address:      Address{Host: "0.0.0.0", Port: 9999},
		Headers:      nil,
		Copyright:    "",
		MaxBodySize:  0,
		ReadTimeout:  0,
		WriteTimeout: 0,
		IdleTimeout:  0,
	}
	s.Init()
	s.Route(GET, "/", func(c *Context) {
		t.Log(c.Request.RequestURI)
		c.String(200, "%s", "Hello World")
	})
	err := s.Listen()
	t.Log(err)
}

func TestNewServerSmooth(t *testing.T) {
	s := Server{
		Debug:        true,
		EnableLog:    true,
		RegSignal:    nil,
		Address:      Address{Host: "0.0.0.0", Port: 9999},
		Headers:      nil,
		Copyright:    "",
		MaxBodySize:  0,
		ReadTimeout:  0,
		WriteTimeout: 0,
		IdleTimeout:  0,
	}
	s.Init()
	s.Route(GET, "/", func(c *Context) {
		t.Log(c.Request.RequestURI)
		c.String(200, "%s", "Hello World")
	})

	s.ListenSmooth()
}

func TestNewServerRegSignal(t *testing.T) {
	s := Server{
		Debug:        true,
		EnableLog:    true,
		RegSignal:    []os.Signal{syscall.SIGTERM},
		Address:      Address{Host: "0.0.0.0", Port: 9999},
		Headers:      nil,
		Copyright:    "",
		MaxBodySize:  0,
		ReadTimeout:  0,
		WriteTimeout: 0,
		IdleTimeout:  0,
	}
	s.Init()
	s.RegSignals(syscall.SIGQUIT)
	s.ListenSmooth()
}

func TestNewServerDuplicateRoute(t *testing.T) {
	s := Server{
		Debug:        true,
		EnableLog:    true,
		RegSignal:    nil,
		Address:      Address{Host: "0.0.0.0", Port: 9999},
		Headers:      nil,
		Copyright:    "",
		MaxBodySize:  0,
		ReadTimeout:  0,
		WriteTimeout: 0,
		IdleTimeout:  0,
	}
	s.Init()
	s.Route(GET, "/", func(c *Context) {
		t.Log(c.Request.RequestURI)
		c.String(200, "%s", "Hello World")
	})
	s.Route(GET, "/", func(c *Context) {
		t.Log(c.Request.RequestURI)
		c.String(200, "%s", "Hello World")
	})
	s.ListenSmooth()
}

func TestNewServerResponse(t *testing.T) {
	s := Server{
		Debug:        true,
		EnableLog:    true,
		RegSignal:    nil,
		Address:      Address{Host: "0.0.0.0", Port: 9999},
		Headers:      nil,
		Copyright:    "",
		MaxBodySize:  0,
		ReadTimeout:  0,
		WriteTimeout: 0,
		IdleTimeout:  0,
	}
	s.Init()
	s.Route(GET, "/1", func(c *Context) {
		t.Log(c.Request.RequestURI)
		c.ResponseStr(200, "HelloWorld")
	})
	s.Route(GET, "/2", func(c *Context) {
		t.Log(c.Request.RequestURI)
		c.ResponseREST(200, map[string]string{"name": "Test"})
	})
	s.Route(GET, "/3", func(c *Context) {
		t.Log(c.Request.RequestURI)
		c.ResponseFile(200, "1.txt")
	})
	s.Route(GET, "/4", func(c *Context) {
		t.Log(c.Request.RequestURI)
		c.ResponseAny("HelloWorld")
	})
	s.Route(GET, "/5", func(c *Context) {
		t.Log(c.Request.RequestURI)
		c.Response(STRING, 200, "HelloWorld")
	})
	s.ListenSmooth()
}

func TestNewServerMiddleware(t *testing.T) {
	s := Server{
		Debug:        false,
		EnableLog:    true,
		RegSignal:    nil,
		Address:      Address{Host: "0.0.0.0", Port: 9999},
		Headers:      nil,
		Copyright:    "",
		MaxBodySize:  0,
		ReadTimeout:  0,
		WriteTimeout: 0,
		IdleTimeout:  0,
	}
	s.Init()
	s.RegMiddle(Wrapper{
		Name:        "middle1",
		WrapperFunc: nil,
	})
	s.RegMiddle(Wrapper{
		Name: "middle2",
		WrapperFunc: func(c *Context) {
			fmt.Println(c.Request.Host)
		},
	})
	s.Route(GET, "/", func(c *Context) {
		t.Log(c.Request.RequestURI)
		c.String(200, "%s", "Hello World")
	})
	s.ListenSmooth()
}
