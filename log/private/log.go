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
)

// fushin内部调用的log日志
// 默认为开发模式打印日志

type ilog struct {
	Prefix string
	Flag   int
	Mode   string
}

var Log ilog
var logMode string

func init() {
	logMode = os.Getenv("FushinMode")
	Log = ilog{
		Prefix: "[Fushin] ",
		Flag:   log.LstdFlags,
		Mode:   logMode,
	}
	Log.init()
}

func (l *ilog) init() {
	log.SetPrefix(l.Prefix)
	log.SetFlags(l.Flag)
}

func (l *ilog) Info(v ...interface{}) {
	if l.Mode == "" || l.Mode == "development" {
		log.Println(append([]interface{}{fmt.Sprintf("%-7s", "[INFO]")}, v...)...)
	}
}

func (l *ilog) InfoF(fmtStr string, v ...interface{}) {
	if l.Mode == "" || l.Mode == "development" {
		log.Print(fmt.Sprintf("%-7s", "[INFO]"), fmt.Sprintf(fmtStr, v...))
	}
}

func (l *ilog) Warn(v ...interface{}) {
	if l.Mode == "" || l.Mode == "development" {
		log.Println(append([]interface{}{fmt.Sprintf("%-7s", "[WARN]")}, v...)...)
	}
}

func (l *ilog) WarnF(fmtStr string, v ...interface{}) {
	if l.Mode == "" || l.Mode == "development" {
		log.Print(fmt.Sprintf("%-7s", "[WARN]"), fmt.Sprintf(fmtStr, v...))
	}
}

func (l *ilog) Error(v ...interface{}) {
	if l.Mode == "" || l.Mode == "development" {
		log.Println(append([]interface{}{fmt.Sprintf("%-7s", "[ERRO]")}, v...)...)
	}
}

func (l *ilog) ErrorF(fmtStr string, v ...interface{}) {
	if l.Mode == "" || l.Mode == "development" {
		log.Print(fmt.Sprintf("%-7s", "[ERRO]"), fmt.Sprintf(fmtStr, v...))
	}
}
