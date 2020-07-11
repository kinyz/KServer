package main

import (
	"KServer/manage"
	"KServer/manage/config"
	"KServer/server/utils"
	"fmt"
	"strconv"
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

	for i := 0; i < 100; i++ {
		fmt.Println("开始发送第", i)
		data := kafka.DataPack().Pack(utils.AgentSendAllClient, 201, "27c340b1-6d1b-4893-a14c-abb1f81829c4", m.Server().GetId(), []byte("全部消息"+strconv.Itoa(i)))

		kafka.Send().Sync("AgentServer_35df2bdb-9ff5-4f12-bb57-8ef8694ba2da", m.Server().GetId(), data)
		//fmt.Println(err)
		//time.Sleep(1 * time.Second)
	}

	select {}
}
