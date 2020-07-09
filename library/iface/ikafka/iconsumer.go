package ikafka

import "github.com/Shopify/sarama"

// 消费者接口
type IConsumer interface {
	NewConsumer(addr []string, offset int64) error
	NewConsumerGroup(addr []string, group string, offset int64) error
	GetConsumer() sarama.Consumer
	GetConsumerGroup() sarama.ConsumerGroup
	CloseConsumer()
	CloseConsumerGroup()
}
