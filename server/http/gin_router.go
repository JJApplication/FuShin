/*
Create: 2022/8/9
Project: FuShin
Github: https://github.com/landers1037
Copyright Renj
*/

// Package http
package http

import (
	"github.com/gin-gonic/gin"
)

func convertWraps(w ...WrapperFunc) []gin.HandlerFunc {
	var h []gin.HandlerFunc
	for _, wf := range w {
		h = append(h, convertWrap(wf))
	}
	return h
}

// 转换为gin的handleFunc
func convertWrap(wrap WrapperFunc) gin.HandlerFunc {
	return func(context *gin.Context) {
		wrap(&Context{context})
	}
}

// 由gin handleFunc转换为wrapFunc
func convertHandle(handler gin.HandlerFunc) WrapperFunc {
	return func(c *Context) {
		handler(c.Context)
	}
}
