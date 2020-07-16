package ikafka

import (
	"KServer/library/kiface/iutils"
	"time"
)

type IKafka interface {
	Send() ISend
	Router() IRouter
}

type IResponse interface {
	GetTopic() string
	GetKey() string
	GetData() iutils.IByte
	GetTimestamp() time.Time
	GetOffset() int64
	GetPartition() int32
}

type BaseResponse interface {
	ResponseHandle(response IResponse)
}

type IRouter interface {
	//	AddRouter()
	// 添加路由
	AddRouter(topic string, response BaseResponse)
	// 添加监听
	StartListen(addr []string, group string, offset int64) func()
	// 添加其他监听
	StartCustomListen(topic []string, addr []string, group string, offset int64) func()
	AddCustomHandle(response BaseResponse)
}

type ISend interface {
	Async(topic string, key string, data []byte)
	Sync(topic string, key string, data []byte) (part int32, offset int64, err error)
	Open(addr []string) error
	Close()
}

type IKafkaConf interface {
	GetAddr() string
}
