/*
Create: 2022/8/14
Project: FuShin
Github: https://github.com/landers1037
Copyright Renj
*/

// Package private
package main

import (
	"github.com/JJApplication/fushin/log"
	"github.com/JJApplication/fushin/server/http"
	"github.com/JJApplication/fushin/server/uds"
)

// Demo 启动一个demo
// demo 包含了server uds, http
// logger of zap
func Demo() {
	// init logger
	logger := log.Logger{
		Name:   "Demo",
		Option: log.DevOption,
		Sync:   true,
	}

	logger.Init()
	// init server
	server := http.Server{
		EnableLog: true,
		Logger:    logger,
		Debug:     false,
		RegSignal: nil,
		Address: http.Address{
			Host: "0.0.0.0",
			Port: 10086,
		},
		Headers: nil,
		PProf:   false,
	}

	// init uds server
	udsServer := uds.UDSServer{
		Name:   "/tmp/DemoUDS",
		Option: uds.DefaultOption,
		Logger: logger,
	}

	// try init all
	server.Init()
	go server.Listen()
	go udsServer.Listen()
	select {}
}
