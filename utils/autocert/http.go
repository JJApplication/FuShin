/*
Create: 2023/3/30
Project: FuShin
Github: https://github.com/landers1037
Copyright Renj
*/

package autocert

import "net"

// 是否可以监听80端口
func canRunHTTP() bool {
	l, err := net.Listen("tcp", ":80")
	if err != nil {
		return false
	}

	l.Close()
	return true
}
