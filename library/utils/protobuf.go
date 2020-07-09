package utils

import (
	"KServer/library/iface/utils"
	"github.com/golang/protobuf/proto"
)

type Protobuf struct {
}

func NewIProtobuf() utils.IProtobuf {
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
