/*
Create: 2023/3/28
Project: FuShin
Github: https://githup.com/landers1037
Copyright Renj
*/

package stream

import "github.com/JJApplication/fushin/utils/buf"

type Parser struct {
	Type int
}

func Parse(t int, data []byte, v interface{}) error {
	f := getTypeParser(t)
	return f(data, v)
}

func MustParse(t int, data []byte, v interface{}) error {
	f := getTypeParser(t)
	return f(data, v)
}

func (p Parser) Parse(data []byte, v interface{}) error {
	f := getTypeParser(p.Type)
	return f(data, v)
}

func (p Parser) MustParse(data []byte, v interface{}) error {
	f := getTypeParser(p.Type)
	return f(data, v)
}

func getTypeParser(t int) func(data []byte, v interface{}) error {
	switch t {
	case JSONType:
		return func(data []byte, v interface{}) error { return JSON.Unmarshal(data, v) }
	case YAMLType:
		return func(data []byte, v interface{}) error {
			return YAML.UnMarshal(data, v)
		}
	case GOBType:
	case FushinBufType:
		// gob需要保证结构体已经被注册进来
		return func(data []byte, v interface{}) error {
			return buf.Decode(data, v)
		}
	case ProtobufType:
		// 暂未支持
		return func(data []byte, v interface{}) error { return JSON.Unmarshal(data, v) }
	default:
		return func(data []byte, v interface{}) error { return JSON.Unmarshal(data, v) }
	}
	return func(data []byte, v interface{}) error { return JSON.Unmarshal(data, v) }
}
