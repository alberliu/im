package main

import (
	"im/config"
	"im/internal/logic/api"
	"im/pkg/db"
	"im/pkg/logger"
	"im/pkg/rpc_cli"
)

func main() {
	db.InitMysql(config.Logic.MySQL)
	db.InitRedis(config.Logic.RedisIP)

	// 初始化RpcClient
	rpc_cli.InitConnIntClient(config.Logic.ConnRPCAddrs)
	rpc_cli.InitUserIntClient(config.Logic.UserRPCAddrs)

	api.StartRpcServer()
	logger.Logger.Info("logic server start")
	select {}
}
