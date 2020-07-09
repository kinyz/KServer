package manage

import (
	"KServer/library/iface/iserver"
	"KServer/library/server"
	"KServer/server/manage/config"
	"KServer/server/manage/pack"
)

type IManage interface {
	Server() iserver.IServer
	Message() pack.IMessage
	DB() pack.IDb
	Client() pack.IClientPack
}
type Manage struct {
	IServer  iserver.IServer
	IMessage pack.IMessage
	Db       pack.IDb
	client   pack.IClientPack
	conf     *config.ManageConfig
}

func NewManage(config *config.ManageConfig) IManage {

	return &Manage{
		IServer:  server.NewIServer(config.Server.Head),
		IMessage: pack.NewIMessagePack(config),
		Db:       pack.NewIDbPack(config),
		client:   pack.NewIClientPack(config),
	}
}
func (m *Manage) Client() pack.IClientPack {
	return m.client
}
func (m *Manage) Server() iserver.IServer {
	return m.IServer
}
func (m *Manage) Message() pack.IMessage {
	return m.IMessage
}
func (m *Manage) DB() pack.IDb {
	return m.Db
}
