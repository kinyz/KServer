package webscoket

import (
	"KServer/library/iface/iwebsocket"
)

type IClient interface {
	Send(data []byte) error
	SendBuff(data []byte) error
	GetConnId() uint32
	Stop()
	GetToken() string
	GetRawConn() iwebsocket.IConnection
}
type Client struct {
	Conn  iwebsocket.IConnection
	Token string
}

func NewClient(conn iwebsocket.IConnection, Token string) IClient {
	return &Client{Conn: conn, Token: Token}
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

func (c *Client) GetRawConn() iwebsocket.IConnection {
	return c.Conn
}
func (c *Client) GetToken() string {
	return c.Token
}
