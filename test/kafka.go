package main

import (
	"KServer/manage"
	"KServer/manage/config"
	pd2 "KServer/manage/discover/pd"
	"KServer/server/utils"
	"KServer/server/utils/msg"
	"fmt"
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

	for i := 0; i < 3; i++ {
		fmt.Println("开始发送第", i)
		//data := kafka.DataPack().Pack(utils.AgentSendAllClient, 201, "27c340b1-6d1b-4893-a14c-abb1f81829c4", m.Server().GetId(), []byte("全部消息"+strconv.Itoa(i)))
		data := &pd2.Discovery{
			Id:       1000,
			Topic:    msg.OauthTopic,
			ServerId: m.Server().GetId(),
			Host:     m.Server().GetHost(),
			Port:     m.Server().GetPort(),
			Type:     "kafka",
		}
		b := m.Message().DataPack().Pack(120, msg.ServiceDiscoveryCheckAllService, "cab02938-4a6a-4e50-b393-94da981e6660",
			m.Server().GetId(), m.Tool().Protobuf().Encode(data))
		kafka.Send().Sync("Client_cab02938-4a6a-4e50-b393-94da981e6660", m.Server().GetId(), b)
		//fmt.Println(err)
		//time.Sleep(1 * time.Second)
	}

	select {}
}
