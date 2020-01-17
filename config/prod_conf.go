package config

import (
	"im/pkg/logger"

	"go.uber.org/zap"
)

func initProdConf() {
	LogicConf = logicConf{
		MySQL:            "root:liu123456@tcp(localhost:3306)/im?charset=utf8&parseTime=true",
		NSQIP:            "127.0.0.1:4150",
		RedisIP:          "127.0.0.1:6379",
		RPCIntListenAddr: ":50000",
		RPCExtListenAddr: ":50001",
		ConnRPCAddrs:     "addrs:///127.0.0.1:60000,127.0.0.1:60001",
		UserRPCAddrs:     "addrs:///127.0.0.1:50300",
	}

	ConnConf = connConf{
		TCPListenAddr: ":8080",
		RPCListenAddr: ":50100",
		LocalAddr:     "127.0.0.1:50100",
		LogicRPCAddrs: "addrs:///127.0.0.1:50000",
	}

	WSConf = wsConf{
		WSListenAddr:  ":8081",
		RPCListenAddr: ":50200",
		LocalAddr:     "127.0.0.1:50200",
		LogicRPCAddrs: "addrs:///127.0.0.1:50000",
	}

	UserConf = userConf{
		MySQL:            "root:liu123456@tcp(localhost:3306)/im?charset=utf8&parseTime=true",
		NSQIP:            "127.0.0.1:4150",
		RedisIP:          "127.0.0.1:6379",
		RPCIntListenAddr: ":50300",
		RPCExtListenAddr: ":50301",
		LogicRPCAddrs:    "addrs:///127.0.0.1:50000",
	}

	logger.Leavel = zap.InfoLevel
	logger.Target = logger.File
}
