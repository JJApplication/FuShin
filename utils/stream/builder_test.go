/*
Create: 2023/3/18
Project: FuShin
Github: https://github.com/landers1037
Copyright Renj
*/

package stream

import "testing"

func TestBuilder(t *testing.T) {
	var types = []int{JSONType, YAMLType, GOBType}
	for _, bt := range types {
		t.Logf("type: %d", bt)
		builder := Builder{Type: bt}
		var data = struct {
			Name string
			Test bool
		}{"Fushin", true}

		res, err := builder.Build(data)
		if err != nil {
			t.Error(err)
			return
		}
		t.Logf("\n%s\n", res)
	}
}
