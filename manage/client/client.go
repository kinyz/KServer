package client

import (
	"KServer/library/iface/isocket"
)

type IClient interface {
	Send(id uint32, msgId uint32, data []byte) error
	SendBuff(id uint32, msgId uint32, data []byte) error
	GetConnId() uint32
	Stop()
	GetToken() string
	GetRawConn() isocket.IConnection
}
type Client struct {
	Conn  isocket.IConnection
	Token string
}

func NewClient(conn isocket.IConnection, Token string) IClient {
	return &Client{Conn: conn, Token: Token}
}
func (c *Client) Send(id uint32, msgId uint32, data []byte) error {
	return c.Conn.SendMsg(id, msgId, data)

}

func (c *Client) SendBuff(id uint32, msgId uint32, data []byte) error {
	return c.Conn.SendBuffMsg(id, msgId, data)
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
