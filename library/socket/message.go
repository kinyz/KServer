package socket

import (
	"KServer/library/iface/isocket"
	"KServer/proto"
)

type Message struct {
	DataLen uint32
	Msg     *proto.Message
}

//创建一个Message消息包
func NewSocketMsgPack() isocket.IMessage {
	return &Message{}
}

func (m *Message) GetId() uint32 {
	return m.Msg.Id
}

func (m *Message) GetMsgId() uint32 {
	return m.Msg.MsgId
}

func (m *Message) GetClientId() string {
	return m.Msg.ClientId
}
func (m *Message) GetServerId() string {
	return m.Msg.ServerId
}
func (m *Message) GetData() []byte {
	return m.Msg.Data
}

func (m *Message) SetData(data []byte) {
	m.Msg.Data = data
}

func (m *Message) SetMessage(message *proto.Message) {
	m.Msg = message
}

func (m *Message) GetDataLen() uint32 {
	return m.DataLen
}
