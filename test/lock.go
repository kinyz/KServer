package main

import (
	"KServer/manage"
	"KServer/manage/config"
	"KServer/server/utils"
	"KServer/server/utils/msg"
	"fmt"
	"time"
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

	fmt.Println(m.Server().GetId())
	q := m.Lock().Lock().Queue("testqueue", 300)
	err := q.Lock(2000, 16000)
	fmt.Println(q.GetTimeOut())
	defer q.UnLock()
	if err != nil {
		fmt.Println(err)
	} else {

		fmt.Println("加锁成功")
	}
	time.Sleep(time.Second * 10)
	//m.Server().Start()

}
