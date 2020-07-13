package iwebsocket

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

	//获取原始数据
	GetRawData() []byte
	// 设置原属数据
	SetRawData(data []byte)
}
