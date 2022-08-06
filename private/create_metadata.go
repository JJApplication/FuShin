/*
Create: 2022/7/7
Project: FuShin
Github: https://github.com/landers1037
Copyright Renj
*/

// Package private
package private

// metadata数据类型 pig为增强的json 支持环境变量和注释

var (
	APPMetaType        = []string{"json", "yaml", "pig", "default"}
	APPMetaTypeDefault = "pig"
	APPProto           = []string{"http", "rpc", "uds"}
	APPProtoDefault    = []string{"http"}
)
