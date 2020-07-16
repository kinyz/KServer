package main

import (
	"KServer/manage"
	"KServer/manage/config"
	"KServer/server/lock/services"
	"KServer/server/utils"
	"KServer/server/utils/msg"
	"fmt"
)

func main() {

	mConf := config.NewManageConfig()

	mConf.DB.Redis = true
	mConf.Server.Head = msg.LockTopic
	mConf.Message.Kafka = true
	mConf.Lock.Open = true
	mConf.Lock.Head = msg.LockTopic
	m := manage.NewManage(mConf)

	// 初始化redisPool
	redisConfig := config.NewRedisConfig(utils.RedisConFile)
	redis := m.DB().Redis()
	if !redis.StartMasterPool(redisConfig.GetMasterAddr(), redisConfig.Master.PassWord, redisConfig.Master.MaxIdle, redisConfig.Master.MaxActive) ||
		!redis.StartSlavePool(redisConfig.GetSlaveAddr(), redisConfig.Slave.PassWord, redisConfig.Slave.MaxIdle, redisConfig.Slave.MaxActive) {
		fmt.Println("Redis 开启失败")
		return
	}

	// 启动消息通道
	kafkaConf := config.NewKafkaConfig(utils.KafkaConFile)
	err := m.Message().Kafka().Send().Open([]string{kafkaConf.GetAddr()})
	if err != nil {
		fmt.Println("消息通道启动失败")
		return
	}

	unLockService := services.NewUnLock(m)

	m.Message().Kafka().AddRouter(msg.LockTopic, msg.LockId, unLockService.UnlockHandle)

	m.Message().Kafka().StartListen([]string{kafkaConf.GetAddr()}, msg.LockTopic, utils.NewOffset)
	m.Server().Start()

}
