package pack

import (
	"KServer/library/iface/redis"
	redis2 "KServer/library/redis"
)

type IDb interface {
	Redis() redis.IRedisPool
}

type Db struct {
	IRedisPool redis.IRedisPool
}

func NewIDbPack() IDb {
	return &Db{IRedisPool: redis2.NewIRedisPool()}
}
func (d *Db) Redis() redis.IRedisPool {
	return d.IRedisPool
}
