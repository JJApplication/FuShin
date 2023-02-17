/*
Create: 2023/2/17
Project: FuShin
Github: https://github.com/landers1037
Copyright Renj
*/

package private

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/gookit/color"
)

type DevLog struct {
	logger *log.Logger
	Color  bool
}

// NewLogger 新建一个内部的DevLog
// 日志记录默认换行
// prefix 日志前缀
// flag log.Flag日志标志
func NewLogger(prefix string, flag int, colored bool) *DevLog {
	return &DevLog{logger: log.New(os.Stdout, prefix, flag), Color: colored}
}

// SetOutput 设置输出的io.writer
func (d *DevLog) SetOutput(w io.Writer) {
	d.logger.SetOutput(w)
}

func (d *DevLog) Info(v ...interface{}) {
	log.Println(append([]interface{}{fmt.Sprintf("%s", d.colored("INFO"))}, v...)...)
}

func (d *DevLog) InfoF(fmtStr string, v ...interface{}) {
	log.Println(fmt.Sprintf("%s %s", d.colored("INFO"), fmt.Sprintf(fmtStr, v...)))
}

func (d *DevLog) Warn(v ...interface{}) {
	log.Println(append([]interface{}{fmt.Sprintf("%s", d.colored("WARN"))}, v...)...)
}

func (d *DevLog) WarnF(fmtStr string, v ...interface{}) {
	log.Println(fmt.Sprintf("%s %s", d.colored("WARN"), fmt.Sprintf(fmtStr, v...)))
}

func (d *DevLog) Error(v ...interface{}) {
	log.Println(append([]interface{}{fmt.Sprintf("%s", d.colored("ERRO"))}, v...)...)
}

func (d *DevLog) ErrorF(fmtStr string, v ...interface{}) {
	log.Println(fmt.Sprintf("%s %s", d.colored("ERRO"), fmt.Sprintf(fmtStr, v...)))
}

func (d *DevLog) Fatal(v ...interface{}) {
	log.Fatalln(append([]interface{}{fmt.Sprintf("%s", d.colored("FATAL"))}, v...)...)
}

func (d *DevLog) FatalF(fmtStr string, v ...interface{}) {
	log.Fatalln(fmt.Sprintf("%s %s", d.colored("FATAL"), fmt.Sprintf(fmtStr, v...)))
}

func (d *DevLog) Print(v ...interface{}) {
	log.Println(v...)
}

func (d *DevLog) PrintF(fmtStr string, v ...interface{}) {
	log.Println(fmt.Sprintf(fmtStr, v...))
}

func (d *DevLog) colored(s string) string {
	var padding = func(s string) string { return fmt.Sprintf(" %s ", s) }
	if d.Color {
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
