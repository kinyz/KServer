package pack

import (
	"KServer/manage/config"
	"KServer/manage/pack/kafkaPack"
)

type IMessage interface {
	Kafka() kafkaPack.IKafkaPack
}

type Message struct {
	IKafkaPack kafkaPack.IKafkaPack
}

func NewIMessagePack(conf *config.ManageConfig) IMessage {
	im := &Message{}
	if conf.Message.Kafka {
		im.IKafkaPack = kafkaPack.NewKafkaPack()
	}
	return im
}

func (m *Message) Kafka() kafkaPack.IKafkaPack {
	return m.IKafkaPack
}
