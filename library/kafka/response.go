package kafka

import (
	"KServer/library/kiface/ikafka"
	"KServer/library/kiface/iutils"
	"time"
)

type Response struct {
	Topic     string
	Key       string
	Data      []byte
	Timestamp time.Time
	Offset    int64
	Partition int32
	iutils.IByte
}

func NewIResponse() ikafka.IResponse {
	return &Response{}
}
func (r *Response) GetTopic() string {
	return r.Topic
}
func (r *Response) GetKey() string {
	return r.Key
}

func (r *Response) GetData() iutils.IByte {
	return r.IByte
}
func (r *Response) GetTimestamp() time.Time {
	return r.Timestamp
}
func (r *Response) GetOffset() int64 {
	return r.Offset
}

func (r *Response) GetPartition() int32 {
	return r.Partition
}
