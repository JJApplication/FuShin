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
	"os"

	"github.com/gookit/color"
)

// fushin内部调用的log日志
// 默认为开发模式打印日志

type ilog struct {
	Prefix string
	Flag   int
	Mode   string
	Color  string
}

var Log ilog
var logMode string

func init() {
	logMode = os.Getenv("FushinMode")
	// 彩色输出
	logColor := os.Getenv("FushinLogColor")
	Log = ilog{
		Prefix: "[Fushin] ",
		Flag:   log.LstdFlags,
		Mode:   logMode,
		Color:  logColor,
	}
	Log.init()
}

func (l *ilog) init() {
	log.SetPrefix(l.Prefix)
	log.SetFlags(l.Flag)
}

func (l *ilog) Info(v ...interface{}) {
	if l.Mode == "" || l.Mode == "development" {
		log.Println(append([]interface{}{fmt.Sprintf("%s ", l.colored("[INFO]"))}, v...)...)
	}
}

func (l *ilog) InfoF(fmtStr string, v ...interface{}) {
	if l.Mode == "" || l.Mode == "development" {
		log.Print(fmt.Sprintf("%s ", l.colored("[INFO]")), fmt.Sprintf(fmtStr, v...))
	}
}

func (l *ilog) Warn(v ...interface{}) {
	if l.Mode == "" || l.Mode == "development" {
		log.Println(append([]interface{}{fmt.Sprintf("%s ", l.colored("[WARN]"))}, v...)...)
	}
}

func (l *ilog) WarnF(fmtStr string, v ...interface{}) {
	if l.Mode == "" || l.Mode == "development" {
		log.Print(fmt.Sprintf("%s ", l.colored("[WARN]")), fmt.Sprintf(fmtStr, v...))
	}
}

func (l *ilog) Error(v ...interface{}) {
	if l.Mode == "" || l.Mode == "development" {
		log.Println(append([]interface{}{fmt.Sprintf("%s ", l.colored("[ERRO]"))}, v...)...)
	}
}

func (l *ilog) ErrorF(fmtStr string, v ...interface{}) {
	if l.Mode == "" || l.Mode == "development" {
		log.Print(fmt.Sprintf("%s ", l.colored("[ERRO]")), fmt.Sprintf(fmtStr, v...))
	}
}

func (l *ilog) colored(s string) string {
	if l.Color == "" || l.Color == "true" || l.Color == "True" {
		switch s {
		case "[INFO]":
			return color.BgBlue.Sprint(s)
		case "[WARN]":
			return color.BgYellow.Sprint(s)
		case "[ERRO]":
			return color.BgRed.Sprint(s)
		default:
			return s
		}
	}
	return s
}
