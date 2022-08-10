/*
Create: 2022/8/9
Project: FuShin
Github: https://github.com/landers1037
Copyright Renj
*/

// Package http
package http

import (
	"github.com/JJApplication/fushin/log/private"
)

func (s *Server) info(v ...interface{}) {
	if s.EnableLog {
		private.Log.Info(v...)
	}
}

func (s *Server) infoF(f string, v ...interface{}) {
	if s.EnableLog {
		private.Log.InfoF(f, v...)
	}
}

func (s *Server) warn(v ...interface{}) {
	if s.EnableLog {
		private.Log.Warn(v...)
	}
}

func (s *Server) warnF(f string, v ...interface{}) {
	if s.EnableLog {
		private.Log.WarnF(f, v...)
	}
}

func (s *Server) error(v ...interface{}) {
	if s.EnableLog {
		private.Log.Error(v...)
	}
}

func (s *Server) errorF(f string, v ...interface{}) {
	if s.EnableLog {
		private.Log.ErrorF(f, v...)
	}
}
