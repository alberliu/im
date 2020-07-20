package main

import (
	"im/config"
	"im/internal/user/api"
	"im/pkg/db"
	"im/pkg/logger"
	"im/pkg/rpc"
)

func main() {
	logger.Init()
	db.InitMysql(config.User.MySQL)
	db.InitRedis(config.User.RedisIP, config.Logic.RedisPassword)

	// 初始化RpcClient
	rpc.InitLogicIntClient(config.User.LogicRPCAddrs)

	api.StartRpcServer()
	logger.Logger.Info("user server start")
	select {}
}
