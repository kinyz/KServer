package proto

import (
	"KServer/library/iface/iutils"
	"KServer/library/utils"
)

type IDataPack interface {
	// pack包
	Pack(id uint32, msgId uint32, clientId string, serverId string, data []byte) []byte
	// unpack包
	UnPack(data []byte) error
	// 获取通信id
	GetId() uint32
	// 获取客户端id
	GetClientId() string
	// 获取服务端id
	GetServerId() string
	// 获取客户端连接id
	GetMsgId() uint32
	// 获取已经unpack到数据
	GetData() iutils.IByte
	// 获取包长度

	GetRawData() []byte
}

type DataPack struct {
	message   *Message
	IProtobuf iutils.IProtobuf
	IByte     iutils.IByte
	RawData   []byte
}

func NewIDataPack() IDataPack {
	return &DataPack{IProtobuf: utils.NewIProtobuf(), IByte: utils.NewIByte()}
}

func (m *DataPack) Pack(id uint32, msgId uint32, clientId string, serverId string, data []byte) []byte {
	//fmt.Println("pack connid =", connId)
	v := &Message{
		Id:       id,
		ClientId: clientId,
		ServerId: serverId,
		MsgId:    msgId,
		Data:     data,
	}
	return m.IProtobuf.Encode(v)
}
func (m *DataPack) UnPack(data []byte) error {

	v := &Message{}
	err := m.IProtobuf.Decode(data, v)
	if err != nil {
		return err
	}
	m.message = v
	m.RawData = data
	//fmt.Println(string(data))
	return nil
}
func (m *DataPack) GetClientId() string {
	return m.message.ClientId
}
func (m *DataPack) GetServerId() string {
	return m.message.ServerId
}
func (m *DataPack) GetMsgId() uint32 {
	return m.message.MsgId
}
func (m *DataPack) GetData() iutils.IByte {
	m.IByte.SetData(m.message.Data)
	return m.IByte
}
func (m *DataPack) GetId() uint32 {
	return m.message.Id
}

func (m *DataPack) GetRawData() []byte {
	return m.RawData
}
