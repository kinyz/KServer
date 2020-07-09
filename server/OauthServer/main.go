package main

import (
	"KServer/library/kafka"
	"KServer/library/redis"
	"KServer/server/OauthServer/services"
	"KServer/server/manage"
	"KServer/server/utils"
)

func main() {
	m := manage.NewManage(utils.OauthTopic)
	// 新建一个服务管理器

	// 启动redis
	redisConf := redis.NewRedisConf(utils.RedisConFile)
	m.DB().Redis().StartMasterPool(redisConf.GetMasterAddr(), redisConf.Master.PassWord, redisConf.Master.MaxIdle, redisConf.Master.MaxActive)
	m.DB().Redis().StartSlavePool(redisConf.GetSlaveAddr(), redisConf.Slave.PassWord, redisConf.Slave.MaxIdle, redisConf.Slave.MaxActive)

	// 启动kafka
	kafkaConf := kafka.NewKafkaConf(utils.KafkaConFile)
	m.Message().Kafka().Send().Open([]string{kafkaConf.GetAddr()})

	oauth := services.NewOauth(m)

	m.Message().Kafka().AddRouter(utils.OauthTopic, utils.OauthMsgId, oauth.ResponseOauth)
	m.Message().Kafka().StartListen([]string{kafkaConf.GetAddr()}, utils.OauthTopic, -1)

	m.Server().Start()
	Close(m)
}

func Close(m manage.IManage) {

}
