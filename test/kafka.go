package main

import (
	"KServer/manage"
	"KServer/manage/config"
	pd2 "KServer/manage/discover/pd"
	"KServer/server/utils"
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
			Topic:    utils.OauthTopic,
			ServerId: m.Server().GetId(),
			Host:     m.Server().GetHost(),
			Port:     m.Server().GetPort(),
			Type:     "kafka",
		}
		b := m.Message().Kafka().DataPack().Pack(utils.ServiceDiscoveryID, utils.ServiceDiscoveryCheckAllService, m.Server().GetId(),
			m.Server().GetId(), m.Tool().Protobuf().Encode(data))
		kafka.Send().Sync(utils.ServiceDiscoveryTopic, m.Server().GetId(), b)
		//fmt.Println(err)
		//time.Sleep(1 * time.Second)
	}

	select {}
}
