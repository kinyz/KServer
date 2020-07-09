package manage

import (
	server2 "KServer/library/iface/server"
	"KServer/library/server"
	"KServer/server/manage/pack"
)

type IManage interface {
	Server() server2.IServer
	Message() pack.IMessage
	DB() pack.IDb
	Client() pack.IClientPack
}
type Manage struct {
	IServer  server2.IServer
	IMessage pack.IMessage
	Db       pack.IDb
	client   pack.IClientPack
}

func NewManage(head string) IManage {
	return &Manage{
		IServer:  server.NewIServer(head),
		IMessage: pack.NewIMessagePack(),
		Db:       pack.NewIDbPack(),
		client:   pack.NewIClientPack(),
	}
}
func (m *Manage) Client() pack.IClientPack {
	return m.client
}

func (m *Manage) Server() server2.IServer {
	return m.IServer
}
func (m *Manage) Message() pack.IMessage {
	return m.IMessage
}
func (m *Manage) DB() pack.IDb {
	return m.Db
}
