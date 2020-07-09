package main

import (
	"KServer/library/iris"
	"KServer/library/mongo"
	"KServer/library/redis"
	"KServer/server/LoginServer/service"
	"KServer/server/manage"
	"KServer/server/utils"
)

func main() {

	m := manage.NewManage(utils.LoginTopic)

	conf := redis.NewRedisConf(utils.RedisConFile)
	m.DB().Redis().StartMasterPool(conf.GetMasterAddr(), conf.Master.PassWord, conf.Master.MaxIdle, conf.Master.MaxActive)
	m.DB().Redis().StartSlavePool(conf.GetSlaveAddr(), conf.Slave.PassWord, conf.Slave.MaxIdle, conf.Slave.MaxActive)

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
