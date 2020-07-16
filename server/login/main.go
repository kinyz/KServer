package main

import (
	"KServer/library/http"
	"KServer/manage"
	"KServer/manage/config"
	"KServer/server/login/service"
	"KServer/server/utils"
	"KServer/server/utils/msg"
	"fmt"
)

func main() {

	ManageConfig := config.NewManageConfig()
	ManageConfig.Server.Head = msg.LoginTopic
	ManageConfig.DB.Redis = true
	ManageConfig.DB.Mongo = true
	ManageConfig.Lock.Open = true
	ManageConfig.Lock.Head = msg.LockTopic
	ManageConfig.Message.Kafka = true

	m := manage.NewManage(ManageConfig)

	kafkaConfig := config.NewKafkaConfig(utils.KafkaConFile)
	err := m.Message().Kafka().Send().Open([]string{kafkaConfig.GetAddr()})
	if err != nil {
		fmt.Println("打开kafkaSend失败:", err)
		return
	}

	redisConfig := config.NewRedisConfig(utils.RedisConFile)
	m.DB().Redis().StartMasterPool(redisConfig.GetMasterAddr(), redisConfig.Master.PassWord, redisConfig.Master.MaxIdle, redisConfig.Master.MaxActive)
	m.DB().Redis().StartSlavePool(redisConfig.GetSlaveAddr(), redisConfig.Slave.PassWord, redisConfig.Slave.MaxIdle, redisConfig.Slave.MaxActive)

	m.DB().Mongo().Start()

	Iris := http.NewIrIrisInterface()
	app := Iris.GetApp()
	user := service.NewUser(m)
	Iris.RegisterPostRouter("/v1/user/accountLogin", user.PreHandler, user.AccountLogin)
	Iris.RegisterPostRouter("/v1/user/accountRegister", user.PreHandler, user.AccountRegister)
	app.Logger().Info("-----------------------------")
	app.Logger().Info("Login Server 启动完毕 ")
	app.Logger().Info("版本: v1.0.0 ")
	app.Logger().Info("作者: moul")
	app.Logger().Info("邮箱: moul@163.com")
	app.Logger().Info("-----------------------------")
	Iris.Init()

}
