package main

import (
	"im/config"
	"im/internal/tcp_conn"
	"im/pkg/rpc_cli"
	"im/pkg/util"
)

func main() {
	// 启动rpc服务
	go func() {
		defer util.RecoverPanic()
		tcp_conn.StartRPCServer()
	}()

	// 初始化Rpc Client
	rpc_cli.InitLogicIntClient(config.ConnConf.LogicRPCAddrs)

	// 启动长链接服务器
	tcp_conn.StartTCPServer()
}
