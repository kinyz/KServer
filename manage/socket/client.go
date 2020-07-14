package socket

import (
	"KServer/library/iface/isocket"
	"KServer/library/socket"
)

type IClient interface {
	/*
		发送消息
		data []byte
		返回值 error
	*/
	Send(data []byte) error
	/*
		发送buff消息
		data []byte
		返回值 error
	*/
	SendBuff(data []byte) error
	/*
		获取Client ConnId
		返回值 uint32
	*/
	GetConnId() uint32
	/*
		停止Client 并执行关闭回调
	*/
	Stop()
	/*
		获取ClientToken
		返回值string
	*/
	GetToken() string
	/*
		获取Client原始Connection
		返回值IConnection
	*/
	GetRawConn() isocket.IConnection
}
type Client struct {
	Conn  isocket.IConnection
	pack  isocket.IDataPack
	Token string
}

func NewClient(conn isocket.IConnection, Token string) IClient {
	return &Client{Conn: conn, Token: Token, pack: socket.NewDataPack()}
}
func (c *Client) Send(data []byte) error {

	return c.Conn.SendMsg(data)

}

func (c *Client) SendBuff(data []byte) error {

	return c.Conn.SendBuffMsg(data)
}

func (c *Client) Stop() {
	c.Conn.Stop()
}

func (c *Client) GetConnId() uint32 {
	return c.Conn.GetConnID()
}

func (c *Client) GetRawConn() isocket.IConnection {
	return c.Conn
}
func (c *Client) GetToken() string {
	return c.Token
}
