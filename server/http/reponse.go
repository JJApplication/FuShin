/*
Create: 2022/8/10
Project: FuShin
Github: https://github.com/landers1037
Copyright Renj
*/

// Package http
package http

import (
	"fmt"

	"github.com/JJApplication/fushin/utils/json"
)

// 注册响应函数
// 可以使用原生的gin响应 也可以使用ctx的响应

const (
	REST   = "REST"
	HTML   = "HTML"
	STRING = "STRING"
	FILE   = "FILE"
)

// Response 响应
func (c *Context) Response(t string, s int, data interface{}) {
	switch t {
	case REST:
		c.JSON(s, data)
		return
	case HTML:
	case FILE:
		c.File(data.(string))
		return
	case STRING:
		c.String(s, "%s", fmt.Sprint(data))
		return
	default:
		c.JSON(s, data)
		return
	}
}

// BeforeResponse 响应前的处理
func (c *Context) BeforeResponse() {

}

// ResponseAny 不做类型推断返回stream
func (c *Context) ResponseAny(data interface{}) {
	d, err := json.Json.Marshal(data)
	if err != nil {
		if _, err = c.Writer.Write(nil); err != nil {
			c.Abort()
		}
		return
	}
	if _, err = c.Writer.Write(d); err != nil {
		c.Abort()
	}
	return
}

// ResponseHtml alias of response with type html
func (c *Context) ResponseHtml(s int, f string) {
	c.Response(HTML, s, f)
	return
}

// ResponseFile alias of response with type file
func (c *Context) ResponseFile(s int, f string) {
	c.Response(FILE, s, f)
	return
}

// ResponseREST alias of response with type rest
func (c *Context) ResponseREST(s int, data interface{}) {
	c.Response(REST, s, data)
	return
}

// ResponseStr alias of response with type string
func (c *Context) ResponseStr(s int, data string) {
	c.Response(STRING, s, data)
	return
}

// ResponseGood 200响应
func (c *Context) ResponseGood(t string, data interface{}) {
	c.Response(t, 200, data)
}

// ResponseBad 错误响应500
func (c *Context) ResponseBad(t string, data interface{}) {
	c.Response(t, 500, data)
}
