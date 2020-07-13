package proto

import "KServer/library/utils"

var proto2 utils.Protobuf

func NewIMessage(id uint32, msgId uint32, clientId string, serverId string, data []byte) []byte {

	return proto2.Encode(&Message{Id: id, MsgId: msgId, ClientId: clientId, ServerId: serverId, Data: data})

}
