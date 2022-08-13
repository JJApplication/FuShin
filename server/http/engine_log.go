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
		if s.Logger == nil {
			private.Log.Info(v...)
		} else {
			s.Logger.Info(v...)
		}
	}
}

func (s *Server) infoF(f string, v ...interface{}) {
	if s.EnableLog {
		if s.Logger == nil {
			private.Log.InfoF(f, v...)
		} else {
			s.Logger.InfoF(f, v...)
		}
	}
}

func (s *Server) warn(v ...interface{}) {
	if s.EnableLog {
		if s.Logger == nil {
			private.Log.Warn(v...)
		} else {
			s.Logger.Warn(v...)
		}
	}
}

func (s *Server) warnF(f string, v ...interface{}) {
	if s.EnableLog {
		if s.Logger == nil {
			private.Log.WarnF(f, v...)
		} else {
			s.Logger.WarnF(f, v...)
		}
	}
}

func (s *Server) error(v ...interface{}) {
	if s.EnableLog {
		if s.Logger == nil {
			private.Log.Error(v...)
		} else {
			s.Logger.Error(v...)
		}
	}
}

func (s *Server) errorF(f string, v ...interface{}) {
	if s.EnableLog {
		if s.Logger == nil {
			private.Log.ErrorF(f, v...)
		} else {
			s.Logger.ErrorF(f, v...)
		}
	}
}
