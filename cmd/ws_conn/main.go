package main

import (
	"im/config"
	"im/internal/ws_conn"
	"im/pkg/rpc_cli"
	"im/pkg/util"
)

func main() {
	// 启动rpc服务
	go func() {
		defer util.RecoverPanic()
		ws_conn.StartRPCServer()
	}()

	// 初始化Rpc Client
	rpc_cli.InitLogicIntClient(config.WSConn.LogicRPCAddrs)

	// 启动长链接服务器
	ws_conn.StartWSServer(config.WSConn.WSListenAddr)
}
