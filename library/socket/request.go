package socket

import (
	"KServer/library/iface/isocket"
)

type Request struct {
	conn isocket.IConnection //已经和客户端建立好的 链接
	msg  isocket.IMessage    //客户端请求的数据
}

//获取请求连接信息
func (r *Request) GetConnection() isocket.IConnection {
	return r.conn
}

//获取请求消息的数据
func (r *Request) GetData() []byte {
	return r.msg.GetData()
}

//获取请求的协议ID
func (r *Request) GetID() uint32 {
	return r.msg.GetId()
}

//获取请求的消息的ID
func (r *Request) GetMsgID() uint32 {
	return r.msg.GetMsgId()
}
