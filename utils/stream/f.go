/*
Create: 2023/2/14
Project: FuShin
Github: https://github.com/landers1037
Copyright Renj
*/

package stream

import "io/ioutil"

// Raw 读取文件内容出错时为空
func Raw(f string) []byte {
	d, e := ioutil.ReadFile(f)
	if e != nil {
		return []byte{}
	}
	return d
}

// RawX 增加错误的Raw
func RawX(f string) ([]byte, error) {
	return ioutil.ReadFile(f)
}
