package config

import (
	"os"
)

var (
	LogicConf logicConf
	ConnConf  connConf
	WSConf    wsConf
	UserConf  userConf
)

// logic配置
type logicConf struct {
	MySQL            string
	NSQIP            string
	RedisIP          string
	RPCIntListenAddr string
	RPCExtListenAddr string
	ConnRPCAddrs     string
	UserRPCAddrs     string
}

// conn配置
type connConf struct {
	TCPListenAddr string
	RPCListenAddr string
	LocalAddr     string
	LogicRPCAddrs string
}

// WS配置
type wsConf struct {
	WSListenAddr  string
	RPCListenAddr string
	LocalAddr     string
	LogicRPCAddrs string
}

// user配置
type userConf struct {
	MySQL            string
	NSQIP            string
	RedisIP          string
	RPCIntListenAddr string
	RPCExtListenAddr string
	LogicRPCAddrs    string
}

func init() {
	env := os.Getenv("im_env")
	switch env {
	case "dev":
		initDevConf()
	case "pre":
		initPreConf()
	case "prod":
		initProdConf()
	default:
		initLocalConf()
	}
}
