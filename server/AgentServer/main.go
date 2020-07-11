package main

import (
	"KServer/manage"
	"KServer/manage/config"
	"KServer/server/AgentServer/response"
	"KServer/server/AgentServer/services"
	"KServer/server/utils"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	mConf := config.NewManageConfig()
	mConf.Socket.Client = true
	mConf.Socket.Server = true
	mConf.DB.Redis = true
	mConf.Server.Head = utils.AgentServerTopic
	mConf.Message.Kafka = true
	// 新建管理器
	m := manage.NewManage(mConf)
	// 管理器启动redis Pool
	redisConf := config.NewRedisConfig(utils.RedisConFile)
	m.DB().Redis().StartMasterPool(redisConf.GetMasterAddr(), redisConf.Master.PassWord, redisConf.Master.MaxIdle, redisConf.Master.MaxActive)
	m.DB().Redis().StartSlavePool(redisConf.GetSlaveAddr(), redisConf.Slave.PassWord, redisConf.Slave.MaxIdle, redisConf.Slave.MaxActive)

	// 启动消息通道
	kafkaConf := config.NewKafkaConfig(utils.KafkaConFile)
	err := m.Message().Kafka().Send().Open([]string{kafkaConf.GetAddr()})
	if err != nil {
		fmt.Println("消息通道启动失败")
		return
	}

	// 启动消息返回
	is := response.NewIServerResponse(m)
	alls := response.NewAllServerResponse(m)

	// 新建socket server
	//socketServer := socket.NewSocket()

	// 注册连接钩子 和连接验证路由
	connect := services.NewConnect(m)
	//注册链接hook回调函数
	m.Socket().Server().SetOnConnStart(connect.DoConnectionBegin)
	m.Socket().Server().SetOnConnStop(connect.DoConnectionLost)
	// 注册socket路由
	m.Socket().Server().AddHandle(utils.OauthId, connect) //添加开始连接路由

	// 添加监听路由
	m.Message().Kafka().AddRouter(m.Server().GetId(), utils.OauthId, connect.ResponseOauth)
	m.Message().Kafka().AddRouter(m.Server().GetId(), utils.AgentSendAllClient, is.SendAllClient) // 通知所有客户端消息

	// 所有服务器接受消息
	//m.Message().Kafka().AddRouter(utils.AgentServerAllTopic, utils.AgentConnStop, alls.RemoveClient) // 通知客户端下线
	m.Message().Kafka().AddRouter(utils.AgentServerAllTopic, utils.AgentAllServerId, alls.ResponseAllServer)

	// 开启监听 和返回通道关闭
	closefunc := m.Message().Kafka().StartListen([]string{kafkaConf.GetAddr()}, m.Server().GetId(), -1)

	//开启scoket服务
	//s.Serve()

	m.Socket().Server().Serve()
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
