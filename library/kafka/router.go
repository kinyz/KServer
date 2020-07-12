package kafka

import (
	"KServer/library/iface/ikafka"
	"KServer/library/iface/iutils"
	"KServer/library/utils"
	"context"
	"fmt"
	"github.com/Shopify/sarama"
	"log"
	"sync"
)

type Router struct {
	Topic        []string
	IConsumer    ikafka.IConsumer
	BaseResponse map[string]ikafka.BaseResponse
	Ready        chan bool
	IByte        iutils.IByte
}

func NewIRouter() ikafka.IRouter {

	return &Router{
		BaseResponse: make(map[string]ikafka.BaseResponse),
		IConsumer:    NewIConsumer(),
		IByte:        utils.NewIByte(),
	}
}

func (r *Router) AddRouter(topic string, response ikafka.BaseResponse) {
	r.BaseResponse[topic] = response

	b := true
	for _, v := range r.Topic {
		if v == topic {
			b = false
			break
		}

	}
	if b {
		r.Topic = append(r.Topic, topic)
		//fmt.Println("新增路由",r.Topic)
	}

}

// Setup is run at the beginning of a new session, before ConsumeClaim
func (r *Router) Setup(sarama.ConsumerGroupSession) error {

	// Mark the consumer as ready
	close(r.Ready)
	return nil
}

// Cleanup is run at the end of a session, once all ConsumeClaim goroutines have exited
func (r *Router) Cleanup(sarama.ConsumerGroupSession) error {

	return nil
}

// ConsumeClaim must start a consumer loop of ConsumerGroupClaim's Messages().
func (r *Router) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		r.IByte.SetData(msg.Value)
		req := &Response{
			Topic:     msg.Topic,
			Key:       string(msg.Key),
			Data:      msg.Value,
			Timestamp: msg.Timestamp,
			Partition: msg.Partition,
			Offset:    msg.Offset,
			IByte:     r.IByte,
		}
		if r.BaseResponse[msg.Topic] != nil {
			r.BaseResponse[msg.Topic].ResponseHandle(req)

		}
		//session.MarkMessage(msg, "")
		//session
		//session.MemberID()
	}
	return nil
}

// 注册组监听
func (r *Router) StartListen(addr []string, group string, offset int64) func() {
	//r.IConsumer =kafka.NewIConsumer()
	err := r.IConsumer.NewConsumerGroup(addr, group, offset)
	if err != nil {
		fmt.Println("[消息监听组]: ", r.Topic, " 启动失败 ", err)
		return nil
	}
	client := r.IConsumer.GetConsumerGroup()
	if client == nil {
		fmt.Println("[消息监听组]: ", r.Topic, " 启动失败")
		return nil
	}
	r.Ready = make(chan bool)
	ctx, cancel := context.WithCancel(context.Background())
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer func() {
			wg.Done()
		}()
		for {
			//fmt.Println(r.Topic)
			//fmt.Println("[消息监听组]: ", []string{"test_go"}, group ,addr,offset)

			if err := client.Consume(ctx, r.Topic, r); err != nil {
				log.Fatalf("Error from consumer: %v", err)
			}
			// check if context was cancelled, signaling that the consumer should stop
			if ctx.Err() != nil {
				log.Println(ctx.Err())
				return
			}
			r.Ready = make(chan bool)
		}
	}()
	<-r.Ready
	fmt.Println("[消息监听组]: ", r.Topic, " 已开启监听")
	return func() {
		cancel()

		wg.Wait()
		if err := client.Close(); err != nil {
			fmt.Println("Error closing client")
		}

		fmt.Println("[消息监听组]: ", r.Topic, " 已关闭监听")

	}
}

// 注册组监听
func (r *Router) StartOtherListen(topic []string, addr []string, group string, offset int64) func() {
	//r.IConsumer =kafka.NewIConsumer()
	err := r.IConsumer.NewConsumerGroup(addr, group, offset)
	if err != nil {
		fmt.Println("[消息监听组]: ", topic, " 启动失败 ", err)
		return nil
	}
	client := r.IConsumer.GetConsumerGroup()
	if client == nil {
		fmt.Println("[消息监听组]: ", topic, " 启动失败")
		return nil
	}
	r.Ready = make(chan bool)
	ctx, cancel := context.WithCancel(context.Background())
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer func() {
			wg.Done()
		}()
		for {
			//fmt.Println(r.Topic)
			//fmt.Println("[消息监听组]: ", []string{"test_go"}, group ,addr,offset)

			if err := client.Consume(ctx, topic, r); err != nil {
				log.Fatalf("Error from consumer: %v", err)
			}
			// check if context was cancelled, signaling that the consumer should stop
			if ctx.Err() != nil {
				log.Println(ctx.Err())
				return
			}
			r.Ready = make(chan bool)
		}
	}()
	<-r.Ready
	fmt.Println("[消息监听组]: ", topic, " 已开启监听")
	return func() {
		cancel()
		wg.Wait()
		for i, v := range r.Topic {
			for _, kv := range topic {
				if v == kv {
					r.Topic = append(r.Topic[:i], r.Topic[i+1:]...)
				}
			}

		}
		if err := client.Close(); err != nil {
			fmt.Println("Error closing client")
		}
		fmt.Println("[消息监听组]: ", topic, " 已关闭监听")

	}
}
