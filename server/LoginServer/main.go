package main

import (
	"KServer/library/iris"
	"KServer/library/mongo"
	"KServer/manage"
	"KServer/manage/config"
	"KServer/server/LoginServer/service"
	"KServer/server/utils"
)

func main() {

	ManageConfig := config.NewManageConfig()
	ManageConfig.Server.Head = utils.LoginTopic
	ManageConfig.DB.Redis = true

	m := manage.NewManage(ManageConfig)

	redisConfig := config.NewRedisConfig(utils.RedisConFile)
	m.DB().Redis().StartMasterPool(redisConfig.GetMasterAddr(), redisConfig.Master.PassWord, redisConfig.Master.MaxIdle, redisConfig.Master.MaxActive)
	m.DB().Redis().StartSlavePool(redisConfig.GetSlaveAddr(), redisConfig.Slave.PassWord, redisConfig.Slave.MaxIdle, redisConfig.Slave.MaxActive)

	Mongo := mongo.NewMongo()
	Mongo.Init()
	Iris := iris.NewIrIrisInterface()
	app := Iris.GetApp()
	user := service.NewUser(m, Mongo)
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
