package kafka

import "KServer/library/iface/ikafka"

type Kafka struct {
	ISend   ikafka.ISend
	IRouter ikafka.IRouter
}

func NewIKafka() ikafka.IKafka {
	return &Kafka{
		ISend:   NewISend(),
		IRouter: NewIRouter(),
	}
}
func (m *Kafka) Send() ikafka.ISend {
	return m.ISend
}

func (m *Kafka) Router() ikafka.IRouter {
	return m.IRouter
}
