/*
Create: 2022/8/20
Project: FuShin
Github: https://github.com/landers1037
Copyright Renj

// 使用New创建一个FushinError对象
// 使用errors.Recover(recover())接收goroutine的recover信息
// 当Recover一个FushinError时不保留堆栈信息 只记录goroutine id
// 使用errors.Panic()创建一个带有FushinError的panic
func TestFushinError() {
	defer func() {
		rec := Recover(recover())
		if rec.HasRecover() {
			t.Log(rec.GetTrace())
			t.Log(rec.GID())
		} else {
			t.Log(rec)
		}
	}()
	// 模拟err
	err := New("this is an error")
	t.Log(err.Error())
	// 模拟panic
	Panic("this is a panic")
	// panic
	panic("panic2")
}
*/

// Package errors
package errors
