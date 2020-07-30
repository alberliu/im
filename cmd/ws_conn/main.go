package main

import (
	"im/config"
	ws_conn2 "im/internal/ws_conn"
	"im/pkg/rpc_cli"
	"im/pkg/util"
)

func main() {
	// 启动rpc服务
	go func() {
		defer util.RecoverPanic()
		ws_conn2.StartRPCServer()
	}()

	// 初始化Rpc Client
	rpc_cli.InitLogicIntClient(config.WSConnConf.LogicRPCAddrs)

	// 启动长链接服务器
	ws_conn2.StartWSServer(config.WSConnConf.WSListenAddr)
}
