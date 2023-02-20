/*
Create: 2022/8/4
Project: FuShin
Github: https://github.com/landers1037
Copyright Renj
*/

// Package private
package private

import (
	"fmt"
	"log"

	"github.com/JJApplication/fushin/pkg"
	"github.com/gookit/color"
)

// fushin内部调用的log日志
// 默认为开发模式打印日志

type ilog struct {
	Prefix string
	Flag   int
	ld     *DevLog
}

var Log ilog

func init() {
	fmt.Println(pkg.Fushin)
	Log = ilog{
		Prefix: fmt.Sprintf("[%s] ", pkg.Fushin),
		Flag:   log.LstdFlags | log.Lshortfile,
	}
	Log.init()
}

func (l *ilog) init() {
	l.ld = NewLogger(l.Prefix, l.Flag, pkg.FushinLogColor)
}

func (l *ilog) Info(v ...interface{}) {
	if pkg.FushinMode == "" || pkg.FushinMode == "development" {
		l.ld.Info(v...)
	}
}

func (l *ilog) InfoF(fmtStr string, v ...interface{}) {
	if pkg.FushinMode == "" || pkg.FushinMode == "development" {
		l.ld.InfoF(fmtStr, v...)
	}
}

func (l *ilog) Warn(v ...interface{}) {
	if pkg.FushinMode == "" || pkg.FushinMode == "development" {
		l.ld.Warn(v...)
	}
}

func (l *ilog) WarnF(fmtStr string, v ...interface{}) {
	if pkg.FushinMode == "" || pkg.FushinMode == "development" {
		l.ld.WarnF(fmtStr, v...)
	}
}

func (l *ilog) Error(v ...interface{}) {
	if pkg.FushinMode == "" || pkg.FushinMode == "development" {
		l.ld.Error(v...)
	}
}

func (l *ilog) ErrorF(fmtStr string, v ...interface{}) {
	if pkg.FushinMode == "" || pkg.FushinMode == "development" {
		l.ld.ErrorF(fmtStr, v...)
	}
}

func (l *ilog) colored(s string) string {
	var padding = func(s string) string { return fmt.Sprintf(" %s ", s) }
	if pkg.FushinLogColor {
		switch s {
		case "INFO":
			return color.BgBlue.Sprint(padding(s))
		case "WARN":
			return color.BgYellow.Sprint(padding(s))
		case "ERRO":
			return color.BgRed.Sprint(padding(s))
		default:
			return padding(s)
		}
	}
	return padding(s)
}
