package kafka

import "KServer/library/iface/kafka"

type Kafka struct {
	ISend   kafka.ISend
	IRouter kafka.IRouter
}

func NewIKafka() kafka.IKafka {
	return &Kafka{
		ISend:   NewISend(),
		IRouter: NewIRouter(),
	}
}
func (m *Kafka) Send() kafka.ISend {
	return m.ISend
}

func (m *Kafka) Router() kafka.IRouter {
	return m.IRouter
}
