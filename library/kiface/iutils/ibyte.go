package iutils

import "github.com/golang/protobuf/proto"

type IByte interface {
	ProtoBuf(value proto.Message) error
	String() string
	Json(value interface{}) error
	Bytes() []byte
	SetData(data []byte)
}
