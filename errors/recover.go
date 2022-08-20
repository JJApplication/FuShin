/*
Create: 2022/8/20
Project: FuShin
Github: https://github.com/landers1037
Copyright Renj
*/

// Package errors
package errors

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/google/uuid"
)

const (
	TraceSIZE = 1 << 20
)

// Recover must pass a recover func
// 当Recover一个FushinError时不保留堆栈信息 只记录goroutine id
func Recover(recover interface{}) FushinError {
	err := FushinError{id: uuid.New().String()}
	if recover != nil {
		// 是否为fushin触发的panic
		if f, ok := recover.(FushinError); ok {
			err.hasRecover = true
			err.stack = fmt.Sprintf("fushinPanic: [%s] %s", f.id, f.message)
			var buf [TraceSIZE]byte
			count := runtime.Stack(buf[:], false)
			id := strings.Fields(strings.TrimPrefix(string(buf[:count]), "goroutine"))[0]
			if id != "" {
				err.gid = id
			}
			return err
		}
		var buf [TraceSIZE]byte
		count := runtime.Stack(buf[:], false)
		err.hasRecover = true
		err.stack = fmt.Sprintf("%s\n%s", recover, buf[:count])
		// 获取goroutine id
		id := strings.Fields(strings.TrimPrefix(string(buf[:count]), "goroutine"))[0]
		if id != "" {
			err.gid = id
		}
	}
	return err
}
