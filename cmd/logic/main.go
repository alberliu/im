package main

import (
	"im/config"
	"im/internal/logic/api"
	"im/pkg/db"
	"im/pkg/logger"
	"im/pkg/rpc"
)

func main() {
	logger.Init()
	db.InitMysql(config.Logic.MySQL)
	db.InitRedis(config.Logic.RedisIP, config.Logic.RedisPassword)

	// 初始化RpcClient
	rpc.InitConnIntClient(config.Logic.ConnRPCAddrs)
	rpc.InitUserIntClient(config.Logic.UserRPCAddrs)

	api.StartRpcServer()
	logger.Logger.Info("logic server start")
	select {}
}
