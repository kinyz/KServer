package kafka

import (
	"KServer/library/kiface/ikafka"
)

type Send struct {
	producer ikafka.IProducer
}

func NewISend() ikafka.ISend {
	return &Send{}

}
func (s *Send) Async(topic string, key string, data []byte) {
	s.producer.SendAsync(topic, key, data)
}
func (s *Send) Sync(topic string, key string, data []byte) (part int32, offset int64, err error) {
	return s.producer.SendSync(topic, key, data)

}
func (s *Send) Open(addr []string) error {
	p := NewIProducer()
	// 初始化生产者
	err := p.NewSyncProducer(addr)
	if err != nil {
		return err
	}
	err = p.NewAsyncProducer(addr)
	if err != nil {
		return err
	}
	s.producer = p
	return nil
}

func (s *Send) Close() {
	s.producer.CloseSyncProducer()
	s.producer.CloseAsyncProducer()
}
