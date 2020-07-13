package pack

import (
	"KServer/manage/config"
	"KServer/manage/pack/kafkaPack"
	"KServer/proto"
)

type IMessage interface {
	Kafka() kafkaPack.IKafkaPack
	DataPack() proto.IDataPack
}

type Message struct {
	IKafkaPack kafkaPack.IKafkaPack
	IDataPack  proto.IDataPack
}

func NewIMessagePack(conf *config.ManageConfig) IMessage {
	im := &Message{IDataPack: proto.NewIDataPack()}
	if conf.Message.Kafka {
		im.IKafkaPack = kafkaPack.NewKafkaPack()
	}
	return im
}

func (m *Message) Kafka() kafkaPack.IKafkaPack {
	return m.IKafkaPack
}

func (m *Message) DataPack() proto.IDataPack {
	return m.IDataPack
}
