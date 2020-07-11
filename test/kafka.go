package main

import (
	"KServer/manage"
	"KServer/manage/config"
	"KServer/server/utils"
	"fmt"
	"time"
)

func main() {
	conf := config.NewManageConfig()
	conf.Message.Kafka = true
	conf.Server.Head = "test"
	m := manage.NewManage(conf)

	kafkaConf := config.NewKafkaConfig(utils.KafkaConFile)
	kafka := m.Message().Kafka()
	//fmt.Println(kafkaConf.GetAddr())
	kafka.Send().Open([]string{kafkaConf.GetAddr()})

	data := kafka.DataPack().Pack(utils.ClientRemove, 201, "27c340b1-6d1b-4893-a14c-abb1f81829c1", m.Server().GetId(), []byte("全部消息"))
	for i := 0; i < 10; i++ {
		fmt.Println("开始发送第", i)
		kafka.Send().Async(utils.ClientTopic+"27c340b1-6d1b-4893-a14c-abb1f81829c1", m.Server().GetId(), data)
		//fmt.Println(err)
		time.Sleep(1 * time.Second)
	}

	select {}
}
