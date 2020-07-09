package main

import (
	"KServer/library/old/message"
	"KServer/server/ChatServer/services"
	"KServer/server/utils"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

// 聊天服务器
func main() {

	msg := message.NewMessage()
	conf := msg.GetKafkaConfig("/conf/kafka.yaml")
	_ = msg.NewConsumerGroup([]string{conf.Host + ":" + conf.Port}, utils.ChatGroup, -2)
	_ = msg.NewConsumer([]string{conf.Host + ":" + conf.Port}, -1)
	chat := services.NewCharHandle()
	f := msg.RegisterGroupListen([]string{utils.ChatTopic}, utils.ChatGroup, chat)
	msg.RegisterListen("message.ChatTopic", 0, -1, chat)
	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)

	select {
	case <-sigterm:
		fmt.Println("terminating: via signal")
		//log.Warnln("terminating: via signal")
	}
	f()
}
