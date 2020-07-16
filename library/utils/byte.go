package utils

import (
	"KServer/library/kiface/iutils"
	"encoding/json"
	"github.com/golang/protobuf/proto"
)

type ByteTool struct {
	*Protobuf
	Data []byte
}

func NewIByte() iutils.IByte {
	return &ByteTool{}
}
func (b *ByteTool) ProtoBuf(value proto.Message) error {
	return b.Protobuf.Decode(b.Data, value)
}
func (b *ByteTool) String() string {
	return string(b.Data)
}
func (b *ByteTool) Json(value interface{}) error {
	return json.Unmarshal(b.Data, value)
}
func (b *ByteTool) Bytes() []byte {
	return b.Data
}
func (b *ByteTool) SetData(data []byte) {
	b.Data = data
}
