package client

import (
	"KServer/library/iface/isocket/ziface"
	"KServer/server/utils/pd"
)

type IClient interface {
	GetUUID() string
	GetAccount() string
	GetToken() string
	GetOauth() bool
	GetPassword() string
	SetOauth(b bool)
	SetConn(conn ziface.IConnection)
	GetConn() ziface.IConnection
	SetAccount(account *pd.Account)
	GetAddr() string
}

type Client struct {
	Account *pd.Account
	Oauth   bool
	Conn    ziface.IConnection
	Addr    string
}

func (c *Client) GetUUID() string {
	return c.Account.GetUUID()
}
func (c *Client) GetAccount() string {
	return c.Account.GetAccount()
}
func (c *Client) SetAccount(account *pd.Account) {

	c.Account = account
}
func (c *Client) GetPassword() string {
	return c.Account.GetPassWord()
}
func (c *Client) GetToken() string {
	return c.Account.GetToken()
}
func (c *Client) GetOauth() bool {
	return c.Oauth
}
func (c *Client) SetOauth(b bool) {
	c.Oauth = b
}

func (c *Client) SetConn(conn ziface.IConnection) {
	c.Addr = conn.RemoteAddr().String()
	c.Conn = conn
}
func (c *Client) GetAddr() string {
	return c.Addr
}

func (c *Client) GetConn() ziface.IConnection {
	return c.Conn
}
