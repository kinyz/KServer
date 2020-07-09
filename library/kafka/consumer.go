package kafka

import (
	"KServer/library/iface/kafka"
	"github.com/Shopify/sarama"
)

type Consumer struct {
	sarama.Consumer
	sarama.ConsumerGroup
}

func NewIConsumer() kafka.IConsumer {
	return &Consumer{}
}

func (c *Consumer) NewConsumer(addr []string, offset int64) error {

	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	config.Version = sarama.V2_3_0_0
	config.Consumer.Offsets.Initial = offset
	consumer, err := sarama.NewConsumer(addr, config)
	if err != nil {
		//fmt.Printf("consumer_test create consumer error %s\n", err.Error())
		return err
	}
	c.Consumer = consumer
	return nil
}
func (c *Consumer) NewConsumerGroup(addr []string, group string, offset int64) error {
	version, err := sarama.ParseKafkaVersion("2.3.0")
	if err != nil {
		return err
	}
	config := sarama.NewConfig()
	config.Version = version
	config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRange // 分区分配策略
	config.Consumer.Offsets.Initial = offset                               // 未找到组消费位移的时候从哪边开始消费
	config.ChannelBufferSize = 2                                           // channel长度
	client, err := sarama.NewConsumerGroup(addr, group, config)
	if err != nil {
		//log.Fatalf("Error creating consumer group client: %v", err)
		return err
	}
	c.ConsumerGroup = client

	return nil
}

func (c *Consumer) GetConsumer() sarama.Consumer {
	return c.Consumer
}

func (c *Consumer) GetConsumerGroup() sarama.ConsumerGroup {
	return c.ConsumerGroup
}

func (c *Consumer) CloseConsumer() {
	_ = c.Consumer.Close()
}
func (c *Consumer) CloseConsumerGroup() {
	_ = c.ConsumerGroup.Close()
}
