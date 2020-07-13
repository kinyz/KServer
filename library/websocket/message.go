package websocket

import (
	"KServer/library/iface/iwebsocket"
	proto2 "KServer/proto"
)

type Message struct {
	Msg     *proto2.Message
	RawData []byte
}

func NewWebSocketMsgPack(msg *proto2.Message) iwebsocket.IMessage {
	return &Message{Msg: msg}
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

func (m *Message) GetRawData() []byte {
	return m.RawData
}

func (m *Message) SetRawData(data []byte) {
	m.RawData = data
}
