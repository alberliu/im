package cache

import (
	"im/pkg/db"
	"im/pkg/util"
)

var RedisUtil = util.NewRedisUtil(db.RedisCli)
