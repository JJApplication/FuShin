/*
Create: 2022/8/21
Project: FuShin
Github: https://github.com/landers1037
Copyright Renj
*/

// Package errors
package errors

// 实现try catch机制

type TryInstance struct {
	innerFunc func()
	isFinal   bool
	exception interface{}
}

// Catch Try实例拥有的Catch方法
// 附带参数exception为出错时包含的错误
func (t *TryInstance) Catch(f func(exception interface{})) *TryInstance {
	if t != nil && t.innerFunc == nil {
		return t
	}

	f(t.exception)
	return t
}

// Finally 总是最后运行
// Finally中执行的函数未被recover()需要保证不会panic
func (t *TryInstance) Finally(f func()) {
	if !t.isFinal {
		t.isFinal = true
		defer f()
	}
}

func Try(f func()) (ti *TryInstance) {
	ti = new(TryInstance)
	ti.innerFunc = f

	// defer 遵循栈规则 FILO
	defer func() {
		r := recover()
		if r != nil {
			ti.exception = r
		}
	}()
	ti.innerFunc()
	return
}
