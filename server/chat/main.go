package main

import (
	"KServer/manage"
	"KServer/manage/config"
	"KServer/server/chat/services"
	"KServer/server/utils"
	"KServer/server/utils/msg"
	"fmt"
	"time"
)

// 聊天服务器
func main() {
	conf := config.NewManageConfig()
	conf.Message.Kafka = true
	conf.Server.Head = msg.ChatTopic

	m := manage.NewManage(conf)
	// 启动消息通道
	kafkaConf := config.NewKafkaConfig(utils.KafkaConFile)
	err := m.Message().Kafka().Send().Open([]string{kafkaConf.GetAddr()})
	if err != nil {
		fmt.Println("消息通道启动失败")
		return
	}
	m.Message().Kafka().Send()

	chat := services.NewChat(m)
	m.Message().Kafka().AddRouter(msg.ChatTopic, msg.ChatId, chat.ResponseChat)
	closeFunc := m.Message().Kafka().StartListen([]string{kafkaConf.GetAddr()}, m.Server().GetId(), utils.NewOffset)

	// 服务中心注册
	m.Message().Kafka().CallRegisterService(msg.ChatId, msg.ChatTopic, m.Server().GetId(), m.Server().GetHost(), m.Server().GetPort(), utils.KafkaType)

	m.Server().Start()

	//服务中心关闭
	m.Message().Kafka().CallLogoutService(msg.ChatId, msg.ChatTopic, m.Server().GetId())

	time.Sleep(5 * time.Second)
	m.Message().Kafka().Send().Close()
	closeFunc()

}
