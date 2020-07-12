package pack

import (
	"KServer/library/iface/imongo"
	"KServer/library/iface/iredis"
	"KServer/library/mongo"
	"KServer/library/redis"
	"KServer/manage/config"
)

type IDb interface {
	Redis() iredis.IRedisPool
	Mongo() imongo.IMongo
}

type Db struct {
	IRedisPool iredis.IRedisPool
	IMongo     imongo.IMongo
}

func NewIDbPack(config *config.ManageConfig) IDb {
	db := &Db{}
	if config.DB.Redis {
		db.IRedisPool = redis.NewIRedisPool()
	}
	if config.DB.Mongo {
		db.IMongo = mongo.NewMongo()
	}
	return db
}
func (d *Db) Redis() iredis.IRedisPool {
	return d.IRedisPool
}

func (d *Db) Mongo() imongo.IMongo {
	return d.IMongo
}
