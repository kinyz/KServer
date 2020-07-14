package main

import (
	"KServer/manage"
	"KServer/manage/config"
	"KServer/server/DiscoveryServer/services"
	"KServer/server/utils"
	"KServer/server/utils/msg"
	"fmt"
)

func main() {

	// 管理器选择开启的服务
	conf := config.NewManageConfig()
	conf.Server.Head = msg.ServiceDiscoveryTopic
	conf.DB.Redis = true
	conf.Message.Kafka = true
	conf.DB.Mongo = true
	m := manage.NewManage(conf)

	// 初始化redisPool
	redisConfig := config.NewRedisConfig(utils.RedisConFile)
	redis := m.DB().Redis()
	if !redis.StartMasterPool(redisConfig.GetMasterAddr(), redisConfig.Master.PassWord, redisConfig.Master.MaxIdle, redisConfig.Master.MaxActive) ||
		!redis.StartSlavePool(redisConfig.GetSlaveAddr(), redisConfig.Slave.PassWord, redisConfig.Slave.MaxIdle, redisConfig.Slave.MaxActive) {
		fmt.Println("Redis 开启失败")
		return
	}

	// 初始化kafka
	kafkaConfig := config.NewKafkaConfig(utils.KafkaConFile)
	kafka := m.Message().Kafka()
	err := kafka.Send().Open([]string{kafkaConfig.GetAddr()})
	if err != nil {
		fmt.Println("Kafka Send 开启失败")
		return
	}

	// 启动mongo数据库
	m.DB().Mongo().Start()

	s := services.NewServiceDiscovery(m)

	kafka.AddRouter(msg.ServiceDiscoveryTopic, msg.ServiceDiscoveryID, s.ServiceHandle)

	kafka.StartListen([]string{kafkaConfig.GetAddr()}, msg.ServiceDiscoveryTopic, utils.NewOffset)

	m.Server().Start()

}
