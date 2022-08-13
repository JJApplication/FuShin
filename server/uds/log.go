/*
Create: 2022/8/5
Project: FuShin
Github: https://github.com/landers1037
Copyright Renj
*/

// Package uds
package uds

import (
	"github.com/JJApplication/fushin/log/private"
)

// Logger 自定义的日志记录器 实现接口即可

func (u *UDSServer) info(v ...interface{}) {
	if !u.Option.LogTrace {
		return
	}
	if u.Logger != nil {
		u.Logger.Info(v...)
	} else {
		private.Log.Info(v...)
	}
}

func (u *UDSServer) infoF(fmt string, v ...interface{}) {
	if !u.Option.LogTrace {
		return
	}
	if u.Logger != nil {
		u.Logger.InfoF(fmt, v...)
	} else {
		private.Log.InfoF(fmt, v...)
	}
}

func (u *UDSServer) warn(v ...interface{}) {
	if !u.Option.LogTrace {
		return
	}
	if u.Logger != nil {
		u.Logger.Warn(v...)
	} else {
		private.Log.Warn(v...)
	}
}

func (u *UDSServer) warnF(fmt string, v ...interface{}) {
	if !u.Option.LogTrace {
		return
	}
	if u.Logger != nil {
		u.Logger.WarnF(fmt, v...)
	} else {
		private.Log.WarnF(fmt, v...)
	}
}

func (u *UDSServer) error(v ...interface{}) {
	if !u.Option.LogTrace {
		return
	}
	if u.Logger != nil {
		u.Logger.Error(v...)
	} else {
		private.Log.Error(v...)
	}
}

func (u *UDSServer) errorF(fmt string, v ...interface{}) {
	if !u.Option.LogTrace {
		return
	}
	if u.Logger != nil {
		u.Logger.ErrorF(fmt, v...)
	} else {
		private.Log.ErrorF(fmt, v...)
	}
}
