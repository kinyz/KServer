package ikafka

import (
	"KServer/library/iface/iutils"
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
	AddRouter(topic string, response BaseResponse)
	StartListen(addr []string, group string, offset int64) func()
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
