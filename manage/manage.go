package manage

import (
	"KServer/library/iface/iserver"
	"KServer/library/server"
	"KServer/manage/config"
	"KServer/manage/pack"
	"KServer/manage/pack/socket"
)

type IManage interface {
	Server() iserver.IServer
	Message() pack.IMessage
	DB() pack.IDb
	Tool() pack.IToolPack
	Socket() socket.ISocketPack
}
type Manage struct {
	IServer    iserver.IServer
	IMessage   pack.IMessage
	Db         pack.IDb
	conf       *config.ManageConfig
	IToolPack  pack.IToolPack
	SocketPack socket.ISocketPack
}

func NewManage(config *config.ManageConfig) IManage {

	return &Manage{
		IServer:    server.NewIServer(config.Server.Head),
		IMessage:   pack.NewIMessagePack(config),
		Db:         pack.NewIDbPack(config),
		IToolPack:  pack.NewIToolPack(),
		SocketPack: socket.NewSocketPack(config),
	}
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
func (m *Manage) Tool() pack.IToolPack {
	return m.IToolPack
}
func (m *Manage) Socket() socket.ISocketPack {
	return m.SocketPack
}
