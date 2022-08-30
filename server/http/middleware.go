/*
Create: 2022/8/10
Project: FuShin
Github: https://github.com/landers1037
Copyright Renj

内置中间件
*/

// Package http
package http

import (
	"fmt"
	"time"

	"github.com/gin-contrib/cache"
	"github.com/gin-contrib/cache/persistence"
	"github.com/gin-contrib/gzip"
)

// 内置中间件

const (
	FushinBuf = "enableFushinBuf"
)

// MiddleWareLogger 路由日志
func MiddleWareLogger() WrapperFunc {
	return convertHandle(ginLogger())
}

// MiddleWareCors 默认的cors
// 允许所有跨域
func MiddleWareCors() WrapperFunc {
	return convertHandle(ginCors())
}

// MiddleWareHeaders 自定义headers
// k-v结构 v为空时清除header
func MiddleWareHeaders(headers map[string]string) WrapperFunc {
	return func(c *Context) {
		// 默认会覆写headers
		for k, v := range headers {
			c.Header(k, v)
		}
	}
}

// MiddleWareGzip 使用gzip
// 默认等级9
func MiddleWareGzip() WrapperFunc {
	return convertHandle(gzip.Gzip(gzip.BestCompression))
}

// MiddleWareGzipLevel 使用gzip 按照等级压缩1-9
// 支持排除某些特殊接口
func MiddleWareGzipLevel(level int, exclude []string) WrapperFunc {
	return convertHandle(gzip.Gzip(level, gzip.WithExcludedPaths(exclude)))
}

var store *persistence.InMemoryStore

func NewStore(duration time.Duration) {
	store = persistence.NewInMemoryStore(duration)
}

// MiddleWareCache 提供基于内存的cache
// 默认的缓存时间为1min 需要使用NewCacheStore创建
// WrapperFunc为要缓存的路由方法
// 更详细的使用请直接使用gin-cache
func MiddleWareCache(duration time.Duration, wrapperFunc WrapperFunc) WrapperFunc {
	if store == nil {
		return func(c *Context) {
			c.Next()
		}
	}
	return convertHandle(cache.CachePage(store, duration, convertWrap(wrapperFunc)))
}

// MiddleWareFushinBuf 开启fushin buf 基于gob的二进制传输
// 仅在header开启enableFushinbuf时 或query?enableFushinbuf=true启用
func MiddleWareFushinBuf() WrapperFunc {
	return func(c *Context) {
		// check headers
		header := c.GetHeader(FushinBuf)
		// check query
		query := c.Query(FushinBuf)
		if header == "true" || query == "true" {
			fmt.Println(header)
			c.Header(FushinBuf, "true")
			c.Set(FushinBuf, true)
		} else {
			c.Next()
		}
	}
}
