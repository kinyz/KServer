package isocket

import "KServer/proto"

/*
	将请求的一个消息封装到message中，定义抽象层接口
*/
type IMessage interface {
	// 获取id
	GetId() uint32

	//获取msgid
	GetMsgId() uint32

	//获取客户端id
	GetClientId() string

	//获取服务器id
	GetServerId() string
	// 获取data
	GetData() []byte
	//设置data
	SetData(data []byte)
	//获取取数据包长度
	GetDataLen() uint32

	SetMessage(message *proto.Message)
}
