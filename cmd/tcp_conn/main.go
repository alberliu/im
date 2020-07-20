package main

import (
	"im/config"
	"im/internal/tcp_conn"
	"im/pkg/logger"
	"im/pkg/rpc"
	"im/pkg/util"
)

func main() {
	logger.Init()
	// 启动rpc服务
	go func() {
		defer util.RecoverPanic()
		tcp_conn.StartRPCServer()
	}()

	// 初始化Rpc Client
	rpc.InitLogicIntClient(config.TCPConn.LogicRPCAddrs)

	// 启动长链接服务器
	tcp_conn.StartTCPServer()
}
