package websocket

import (
	"KServer/library/iface/iwebsocket"
)

type Request struct {
	conn    iwebsocket.IConnection //已经和客户端建立好的 链接
	message iwebsocket.IMessage    //客户端请求的数据
}

//获取请求连接信息
func (r *Request) GetConnection() iwebsocket.IConnection {
	return r.conn
}

//获取请求消息的数据
func (r *Request) GetMessage() iwebsocket.IMessage {
	return r.message
}
