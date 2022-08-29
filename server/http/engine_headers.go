/*
Create: 2022/8/29
Project: FuShin
Github: https://github.com/landers1037
Copyright Renj
*/

// Package http
package http

// 响应头

func (s *Server) SetHeaders(headers map[string]string) {
	s.engine.Use(convertWrap(MiddleWareHeaders(headers)))
}

// 添加版本信息响应头
func (s *Server) addCopyright() {
	if s.Copyright != "" {
		s.engine.Use(convertWrap(MiddleWareHeaders(map[string]string{"Copyright": s.Copyright})))
	}
}

// 添加预置headers
func (s *Server) addPresetHeaders() {
	s.SetHeaders(s.Headers)
}
