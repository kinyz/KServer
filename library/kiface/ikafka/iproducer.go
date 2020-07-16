package ikafka

import "github.com/Shopify/sarama"

// 生产者接口
type IProducer interface {
	NewSyncProducer(addr []string) error
	NewAsyncProducer(addr []string) error
	SendSync(topic string, key string, data []byte) (part int32, offset int64, err error)
	SendAsync(topic string, key string, data []byte)
	GetSyncProducer() sarama.SyncProducer
	GetAsyncProducer() sarama.AsyncProducer
	CloseAsyncProducer()
	CloseSyncProducer()
}
