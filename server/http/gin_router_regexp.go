/*
Create: 2023/2/14
Project: FuShin
Github: https://github.com/landers1037
Copyright Renj
*/

package http

import (
	"github.com/jxskiss/ginregex"
)

// RouteRegex 使用正则匹配路由
func (s *Server) RouteRegex(method, uri string, wrap ...WrapperFunc) {
	ginregex.NewMatcher(method, uri, convertWraps(wrap...)...)
}

// PatchServerWithRegex 对Server进行封装
// 后续所有的路由定义将使用正则匹配
func PatchServerWithRegex(s *Server) *ginregex.RegexRouter {
	return ginregex.New(s.engine, nil)
}
