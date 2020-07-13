package kafka

import (
	"KServer/library/iface/ikafka"
	"fmt"
	"github.com/Shopify/sarama"
	"log"
	"time"
)

type Producer struct {
	sarama.AsyncProducer
	sarama.SyncProducer
}

func NewIProducer() ikafka.IProducer {
	return &Producer{}
}

func (p *Producer) NewSyncProducer(addr []string) error {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Timeout = 5 * time.Second
	config.Version = sarama.V2_3_0_0
	sync, err := sarama.NewSyncProducer(addr, config)
	if err != nil {
		//fmt.Println("[ERROR]:NewSyncProducer fail! err=" + err.Error())
		return err
	}
	p.SyncProducer = sync
	return nil
}
func (p *Producer) NewAsyncProducer(addr []string) error {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = false
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Timeout = 5 * time.Second
	config.Version = sarama.V2_3_0_0
	async, err := sarama.NewAsyncProducer(addr, config)
	if err != nil {
		//	fmt.Println("[ERROR]:NewAsyncProducer fail! err=" + err.Error())
		return err
	}
	p.AsyncProducer = async
	return nil
}

// AsyncSendMsg 同步生产者
// 返回 part, offset, err
func (p *Producer) SendSync(topic string, key string, data []byte) (int32, int64, error) {

	msg := &sarama.ProducerMessage{
		Topic:     topic,
		Key:       sarama.StringEncoder(key),
		Value:     sarama.ByteEncoder(data),
		Timestamp: time.Time{},
	}
	part, offset, err := p.SyncProducer.SendMessage(msg)
	if err != nil {
		log.Printf("send message(%s) err=%s \n", data, err)
		return 0, 0, err
	} else {

		return part, offset, err
	}
}

// AsyncSendMsg 异步生产者
// 并发量大时，必须采用这种方式
func (p *Producer) SendAsync(topic string, key string, data []byte) {
	async := p.GetAsyncProducer()
	go func(as sarama.AsyncProducer) {
		errors := as.Errors()
		success := as.Successes()
		for {
			select {
			case err := <-errors:
				if err != nil {
					fmt.Println(err)
				}
			case <-success:
			}
		}
	}(async)

	msg := &sarama.ProducerMessage{
		Topic:     topic,
		Key:       sarama.StringEncoder(key),
		Value:     sarama.ByteEncoder(data),
		Timestamp: time.Time{},
	}
	async.Input() <- msg
}

func (p *Producer) GetSyncProducer() sarama.SyncProducer {
	return p.SyncProducer
}
func (p *Producer) GetAsyncProducer() sarama.AsyncProducer {
	return p.AsyncProducer
}
func (p *Producer) CloseAsyncProducer() {
	if err := p.AsyncProducer.Close(); err != nil {
		fmt.Println("CloseAsyncProducer Fail err=", err)
	}

}
func (p *Producer) CloseSyncProducer() {
	if err := p.SyncProducer.Close(); err != nil {
		fmt.Println("CloseSyncProducer Fail err=", err)
	}
}
