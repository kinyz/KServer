package socket

import (
	"KServer/library/iface/isocket"
	"KServer/library/socket"
)

type IClient interface {
	Send(data []byte) error
	SendBuff(data []byte) error
	GetConnId() uint32
	Stop()
	GetToken() string
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
