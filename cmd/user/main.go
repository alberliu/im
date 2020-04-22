package main

import (
	"im/config"
	"im/internal/user/api"
	"im/pkg/db"
	"im/pkg/logger"
	"im/pkg/rpc_cli"
)

func main() {
	db.InitMysql(config.User.MySQL)
	db.InitRedis(config.User.RedisIP)

	// 初始化RpcClient
	rpc_cli.InitLogicIntClient(config.User.LogicRPCAddrs)

	api.StartRpcServer()
	logger.Logger.Info("user server start")
	select {}
}
