package pack

type IMessage interface {
	Kafka() IKafkaPack
}

type Message struct {
	IKafkaPack IKafkaPack
}

func NewIMessagePack() IMessage {
	return &Message{IKafkaPack: NewKafkaPack()}
}

func (m *Message) Kafka() IKafkaPack {
	return m.IKafkaPack
}
