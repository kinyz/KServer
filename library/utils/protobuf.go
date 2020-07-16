package utils

import (
	"KServer/library/kiface/iutils"
	"github.com/golang/protobuf/proto"
)

type Protobuf struct {
}

func NewIProtobuf() iutils.IProtobuf {
	return &Protobuf{}
}
func (p *Protobuf) Encode(table proto.Message) []byte {
	data, err := proto.Marshal(table)
	if err != nil {
		return nil
	}
	return data
}

func (p *Protobuf) Decode(b []byte, m proto.Message) error {
	return proto.Unmarshal(b, m)
}
