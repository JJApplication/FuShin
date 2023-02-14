/*
Create: 2023/2/14
Project: FuShin
Github: https://github.com/landers1037
Copyright Renj
*/

package stream

import (
	"gopkg.in/yaml.v3"
)

type ymlStream struct{}

var YAML = ymlStream{}
var YML = YAML

func (y ymlStream) Marshal(in interface{}) (out []byte, err error) {
	return yaml.Marshal(in)
}

func (y ymlStream) UnMarshal(data []byte, out interface{}) error {
	return yaml.Unmarshal(data, out)
}
