package iredis

import "github.com/golang/protobuf/proto"

//set value接口
type ISetValue interface {
	ProtoBuf(value proto.Message) (reply interface{}, err error)
	String(value string) (reply interface{}, err error)
	Json(value interface{}) (reply interface{}, err error)
	Bytes(value []byte) (reply interface{}, err error)
	Value(value interface{}) (reply interface{}, err error)
}

// get value接口
type IGetValue interface {
	ProtoBuf(value proto.Message) error
	Json(value interface{}) error
	String() string
	Bytes() []byte
	Value() (reply interface{}, err error)
}

// value 接口
type IValue interface {
	Get(key string) IGetValue
	Set(key string) ISetValue
	Do(commandName string, args ...interface{}) (reply interface{}, err error)
	Check(key string) bool
	Del(key string) (reply interface{}, err error)
}
