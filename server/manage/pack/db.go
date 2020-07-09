package pack

import (
	"KServer/library/iface/iredis"
	"KServer/library/redis"
	"KServer/server/manage/config"
)

type IDb interface {
	Redis() iredis.IRedisPool
}

type Db struct {
	IRedisPool iredis.IRedisPool
}

func NewIDbPack(config *config.ManageConfig) IDb {
	db := &Db{}
	if config.DB.Redis {
		db.IRedisPool = redis.NewIRedisPool()
	}
	return db
}
func (d *Db) Redis() iredis.IRedisPool {
	return d.IRedisPool
}
