/*
Create: 2023/2/17
Project: FuShin
Github: https://github.com/landers1037
Copyright Renj
*/

package module

// 在主程序中注册用于从module中获取信息

// Module 包含所注册模块的全部信息
type Module struct {
	Name   string
	Status bool
}

// M 插件必须实现的接口
type M interface {
	Name() string
	Hooks(s interface{})
	Enable()
	Disable()
	Extra() map[string]interface{}
}

type Modules map[string]Module

func CreateModules() Modules {
	m := make(Modules, 0)
	return m
}

func AddModule(ms Modules, m Module) {
	ms[m.Name] = m
}

func RmModule(ms Modules, m Module) {
	if _, ok := ms[m.Name]; ok {
		delete(ms, m.Name)
	}
}

func (m *Module) Enable() {
	m.Status = true
}

func (m *Module) Disable() {
	m.Status = false
}
