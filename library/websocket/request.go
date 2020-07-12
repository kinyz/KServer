package websocket

import (
	"KServer/library/iface/iwebsocket"
)

type Request struct {
	conn iwebsocket.IConnection //已经和客户端建立好的 链接
	data []byte                 //客户端请求的数据
	id   uint32
}

//获取请求连接信息
func (r *Request) GetConnection() iwebsocket.IConnection {
	return r.conn
}

//获取请求消息的数据
func (r *Request) GetData() []byte {
	return r.data
}

//获取请求消息的数据
func (r *Request) GetId() uint32 {
	return r.id
}
