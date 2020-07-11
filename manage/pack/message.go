package pack

import "KServer/manage/config"

type IMessage interface {
	Kafka() IKafkaPack
}

type Message struct {
	IKafkaPack IKafkaPack
}

func NewIMessagePack(conf *config.ManageConfig) IMessage {
	im := &Message{}
	if conf.Message.Kafka {
		im.IKafkaPack = NewKafkaPack()
	}
	return im
}

func (m *Message) Kafka() IKafkaPack {
	return m.IKafkaPack
}
