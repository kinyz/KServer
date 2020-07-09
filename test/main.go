package main

import (
	"KServer/library/old/message"
	"KServer/server/utils"
	"fmt"
	"strconv"
)

func main() {

	msg := message.NewMessage()
	conf := msg.GetKafkaConfig("/conf/kafka.yaml")
	err := msg.NewAsyncProducer([]string{conf.Host + ":" + conf.Port})
	if err != nil {
		fmt.Println("创建AsyncProducer失败")
	}
	for i := 0; i < 2; i++ {
		_ = msg.SendAsyncMessage(utils.ChatTopic, utils.ChatSendKey, utils.ChatSendWord, "异步世界说话"+strconv.Itoa(i))
		_ = msg.SendAsyncMessage(utils.ChatTopic, utils.ChatSendKey, utils.ChatSendPrivate, "异步私人说话"+strconv.Itoa(i))
		//time.Sleep(5 * time.Second)
	}

	err2 := msg.NewSyncProducer([]string{conf.Host + ":" + conf.Port})
	if err2 != nil {
		fmt.Println("创建SyncProducer失败")
	}
	part, offset, err3 := msg.SendSyncMessage(utils.ChatTopic, utils.ChatSendKey, utils.ChatSendWord, "同步世界说话")

	if err3 != nil {
		fmt.Println("同步发送失败", err)
	}
	fmt.Println(part, offset)

	msg.CloseSyncProducer()
	msg.CloseAsyncProducer()
	select {}
}
