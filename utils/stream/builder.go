/*
Create: 2023/3/17
Project: FuShin
Github: https://github.com/landers1037
Copyright Renj
*/

package stream

import "github.com/JJApplication/fushin/utils/buf"

// Builder 生成器 基于当前流类型构造数据流
type Builder struct {
	Type int
}

const (
	JSONType = iota
	YAMLType
	GOBType
	FushinBufType
	ProtobufType
)

func Build(t int, v interface{}) ([]byte, error) {
	f := getTypeBuilder(t)
	return f(v)
}

func MustBuild(t int, v interface{}) []byte {
	f := getTypeBuilder(t)
	d, _ := f(v)
	return d
}

func (b Builder) Build(v interface{}) ([]byte, error) {
	f := getTypeBuilder(b.Type)
	return f(v)
}

func (b Builder) MustBuild(v interface{}) []byte {
	f := getTypeBuilder(b.Type)
	d, _ := f(v)
	return d
}

func getTypeBuilder(t int) func(v interface{}) ([]byte, error) {
	switch t {
	case JSONType:
		return func(v interface{}) ([]byte, error) { return JSON.Marshal(v) }
	case YAMLType:
		return func(v interface{}) ([]byte, error) {
			return YAML.Marshal(v)
		}
	case GOBType:
	case FushinBufType:
		// gob需要保证结构体已经被注册进来
		return func(v interface{}) ([]byte, error) {
			return buf.Encode(v)
		}
	case ProtobufType:
		// 暂未支持
		return func(v interface{}) ([]byte, error) { return JSON.Marshal(v) }
	default:
		return func(v interface{}) ([]byte, error) { return JSON.Marshal(v) }
	}
	return func(v interface{}) ([]byte, error) { return JSON.Marshal(v) }
}
