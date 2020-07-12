package main

import (
	"KServer/manage"
	"KServer/manage/config"
	"KServer/server/OauthServer/services"
	"KServer/server/utils"
	"time"
)

func main() {
	conf := config.NewManageConfig()
	conf.Message.Kafka = true
	conf.Server.Head = utils.OauthTopic
	conf.DB.Redis = true
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

	m.Message().Kafka().AddRouter(utils.OauthTopic, utils.OauthId, oauth.ResponseOauth)
	m.Message().Kafka().StartListen([]string{kafkaConf.GetAddr()}, utils.OauthTopic, -1)

	// 服务中心注册服务
	//d:=generalService.NewIDiscovery(m)
	m.Message().Kafka().CallRegisterService(utils.OauthId, utils.OauthTopic, m.Server().GetId(), m.Server().GetHost(), m.Server().GetPort(), utils.KafkaType)

	m.Server().Start()

	m.Message().Kafka().CallLogoutService(utils.OauthId, utils.OauthTopic, m.Server().GetId())
	time.Sleep(5 * time.Second)
	Close(m)

}

func Close(m manage.IManage) {
	// 注销服务中心

	_ = m.DB().Redis().CloseMaster()
	_ = m.DB().Redis().CloseSlave()

}
