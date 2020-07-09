package utils

import "github.com/golang/protobuf/proto"

type IProtobuf interface {
	Encode(table proto.Message) []byte
	Decode(b []byte, m proto.Message) error
}
