package main

import (
	"im/config"
	"im/internal/logic/api"
	"im/pkg/db"
	"im/pkg/logger"
	"im/pkg/rpc_cli"
)

func main() {
	db.InitMysql(config.LogicConf.MySQL)
	db.InitRedis(config.LogicConf.RedisIP)

	// 初始化RpcClient
	rpc_cli.InitConnIntClient(config.LogicConf.ConnRPCAddrs)
	rpc_cli.InitUserIntClient(config.LogicConf.UserRPCAddrs)

	api.StartRpcServer()
	logger.Logger.Info("logic server start")
	select {}
}
