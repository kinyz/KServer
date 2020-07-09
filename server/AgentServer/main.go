package main

import (
	"KServer/library/socket/znet"
	"KServer/server/AgentServer/response"
	"KServer/server/AgentServer/services"
	"KServer/server/manage"
	"KServer/server/manage/config"
	"KServer/server/utils"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	mConf := config.NewManageConfig()
	mConf.Client = true
	mConf.DB.Redis = true
	mConf.Server.Head = utils.AgentServerTopic
	mConf.Message.Kafka = true
	// 新建管理器
	m := manage.NewManage(mConf)
	// 管理器启动redis Pool
	Redisconf := config.NewRedisConfig(utils.RedisConFile)
	m.DB().Redis().StartMasterPool(Redisconf.GetMasterAddr(), Redisconf.Master.PassWord, Redisconf.Master.MaxIdle, Redisconf.Master.MaxActive)
	m.DB().Redis().StartSlavePool(Redisconf.GetSlaveAddr(), Redisconf.Slave.PassWord, Redisconf.Slave.MaxIdle, Redisconf.Slave.MaxActive)

	// 启动消息通道
	conf2 := config.NewKafkaConfig(utils.KafkaConFile)
	err := m.Message().Kafka().Send().Open([]string{conf2.GetAddr()})
	if err != nil {
		fmt.Println("消息通道启动失败")
		return
	}

	// 启动消息返回
	is := response.NewIServerResponse(m)
	alls := response.NewAllServerResponse(m)

	// 新建socket server
	socket := znet.NewServer()

	// 注册连接钩子 和连接验证路由
	connect := services.NewConnect(m)
	//注册链接hook回调函数
	socket.SetOnConnStart(connect.DoConnectionBegin)
	socket.SetOnConnStop(connect.DoConnectionLost)
	socket.AddRouter(utils.OauthMsgId, connect)

	// 添加监听路由
	m.Message().Kafka().AddRouter(m.Server().GetId(), utils.OauthMsgId, connect.ResponseOauth)
	m.Message().Kafka().AddRouter(m.Server().GetId(), utils.AgentConnStop, is.ResponseOauth)
	m.Message().Kafka().AddRouter(utils.AgentServerAllTopic, utils.AgentAllConnStop, alls.ResponseAllStop)

	// 开启监听 和返回通道关闭
	closefunc := m.Message().Kafka().StartListen([]string{conf2.GetAddr()}, m.Server().GetId(), -1)

	//开启scoket服务
	//s.Serve()

	socket.Start()
	fmt.Println("[服务器加载完毕]")

	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	signal.Notify(sigs, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		<-sigs
		//fmt.Println()
		//fmt.Println(sig)
		done <- true
	}()

	//fmt.Println("awaiting signal")

	<-done

	fmt.Println("Server Close...")
	// 关闭消息监听

	// 关闭socket
	//socket.Stop()
	// 关闭redis
	_ = m.DB().Redis().CloseMaster()
	_ = m.DB().Redis().CloseSlave()
	// 关闭消息通道
	m.Message().Kafka().Send().Close()
	closefunc()

}
