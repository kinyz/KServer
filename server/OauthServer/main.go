package main

import (
	"KServer/server/OauthServer/services"
	"KServer/server/manage"
	"KServer/server/manage/config"
	"KServer/server/utils"
)

func main() {
	conf := config.NewManageConfig()
	conf.Message.Kafka = true
	conf.Server.Head = utils.OauthTopic
	m := manage.NewManage(conf)
	// 新建一个服务管理器

	// 启动redis
	redisConf := config.NewRedisConfig(utils.RedisConFile)
	m.DB().Redis().StartMasterPool(redisConf.GetMasterAddr(), redisConf.Master.PassWord, redisConf.Master.MaxIdle, redisConf.Master.MaxActive)
	m.DB().Redis().StartSlavePool(redisConf.GetSlaveAddr(), redisConf.Slave.PassWord, redisConf.Slave.MaxIdle, redisConf.Slave.MaxActive)

	// 启动kafka
	kafkaConf := config.NewKafkaConfig(utils.KafkaConFile)
	_ = m.Message().Kafka().Send().Open([]string{kafkaConf.GetAddr()})

	oauth := services.NewOauth(m)

	m.Message().Kafka().AddRouter(utils.OauthTopic, utils.OauthMsgId, oauth.ResponseOauth)
	m.Message().Kafka().StartListen([]string{kafkaConf.GetAddr()}, utils.OauthTopic, -1)

	m.Server().Start()
	Close(m)
}

func Close(m manage.IManage) {

	_ = m.DB().Redis().CloseMaster()
	_ = m.DB().Redis().CloseSlave()

}
