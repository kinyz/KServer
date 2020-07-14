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
	ready        chan bool
	IByte        iutils.IByte
	CustomHandle ikafka.BaseResponse
}

func NewIRouter() ikafka.IRouter {

	return &Router{
		BaseResponse: make(map[string]ikafka.BaseResponse),
		IConsumer:    NewIConsumer(),
		IByte:        utils.NewIByte(),
		ready:        make(chan bool),
	}
}
func (r *Router) AddCustomHandle(response ikafka.BaseResponse) {

	if r.CustomHandle != nil {
		return
	}
	r.CustomHandle = response
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

	//	fmt.Println("111")

	close(r.ready)
	return nil
}

// Cleanup is run at the end of a session, once all ConsumeClaim goroutines have exited
func (r *Router) Cleanup(sarama.ConsumerGroupSession) error {

	return nil
}

// ConsumeClaim must start a consumer loop of ConsumerGroupClaim's Messages().
func (r *Router) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {

		//fmt.Println("ConsumeClaim",string(msg.Value))
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
			//print("执行1")
			r.BaseResponse[msg.Topic].ResponseHandle(req)
			//break
		} else {
			if r.CustomHandle != nil {
				//	print("执行2")
				r.CustomHandle.ResponseHandle(req)
			}
		}
		session.MarkMessage(msg, "")

	}
	return nil
}

// 注册组监听
func (r *Router) StartListen(addr []string, group string, offset int64) func() {
	//r.IConsumer =kafka.NewIConsumer()

	config := sarama.NewConfig()
	config.Version = sarama.V2_3_0_0
	config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRange // 分区分配策略
	config.Consumer.Offsets.Initial = offset                               // 未找到组消费位移的时候从哪边开始消费
	config.ChannelBufferSize = 2                                           // channel长度
	client, err := sarama.NewConsumerGroup(addr, group, config)
	//client := r.IConsumer.GetConsumerGroup()
	if err != nil {
		fmt.Println("[消息监听组]: ", r.Topic, " 启动失败")
	}

	//r.Ready = make(chan bool)
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
			r.ready = make(chan bool)
		}
	}()
	<-r.ready
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
func (r *Router) StartCustomListen(topic []string, addr []string, group string, offset int64) func() {
	config := sarama.NewConfig()
	config.Version = sarama.V2_3_0_0
	config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRange // 分区分配策略
	config.Consumer.Offsets.Initial = offset                               // 未找到组消费位移的时候从哪边开始消费
	config.ChannelBufferSize = 2                                           // channel长度
	client, err := sarama.NewConsumerGroup(addr, group, config)

	if err != nil {
		fmt.Println("[消息监听组]: ", r.Topic, " 启动失败")
	}
	//r.Ready = make(chan bool)
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
			//		r.Ready = make(chan bool)
		}
	}()
	//<-r.Ready
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
