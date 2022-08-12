/*
Create: 2022/8/11
Project: FuShin
Github: https://github.com/landers1037
Copyright Renj
*/

// Package env
package env

import (
	"fmt"
	"os"
	"reflect"
	"sync"
	"testing"
)

func TestEnvLoader(t *testing.T) {
	os.Setenv("str", "test")
	os.Setenv("int", "1")
	os.Setenv("bool", "true")
	os.Setenv("float", "1.0")
	os.Setenv("any", "any")
	load := EnvLoader{}
	t.Log(load.Get("str").Raw(), reflect.TypeOf(load.Get("str").Raw()))
	t.Log(load.Get("int").Raw(), reflect.TypeOf(load.Get("int").Int()))
	t.Log(load.Get("bool").Raw(), reflect.TypeOf(load.Get("bool").Bool()))
	t.Log(load.Get("float").Raw(), reflect.TypeOf(load.Get("float").Float64()))
	t.Log(load.Get("any").Raw(), reflect.TypeOf(load.Get("any").Interface()))
}

func TestEnvLoaderConcurrent(t *testing.T) {
	wg := sync.WaitGroup{}
	load := EnvLoader{}
	for i := range [100]struct{}{} {
		wg.Add(1)
		os.Setenv(fmt.Sprintf("str%d", i), "test")
		t.Log(load.Get(fmt.Sprintf("str%d", i)).Raw(), reflect.TypeOf(load.Get(fmt.Sprintf("str%d", i)).Raw()))
		wg.Done()
	}

	wg.Wait()
}
