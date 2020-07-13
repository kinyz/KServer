package manage

import (
	"KServer/library/iface/iserver"
	"KServer/library/server"
	"KServer/manage/config"
	"KServer/manage/discover"
	"KServer/manage/pack"
	"KServer/manage/pack/socket"
	"KServer/manage/pack/websocket"
)

type IManage interface {
	// 服务器管理
	Server() iserver.IServer
	// 通信协议
	Message() pack.IMessage
	// 数据库
	DB() pack.IDb
	// 工具
	Tool() pack.IToolPack
	// Socket
	Socket() socket.ISocketPack
	// WebSocket
	WebSocket() websocket.IWebSocketPack
	// 服务管理器
	Discover() discover.IDiscover
}
type Manage struct {
	IServer       iserver.IServer
	IMessage      pack.IMessage
	Db            pack.IDb
	conf          *config.ManageConfig
	IToolPack     pack.IToolPack
	SocketPack    socket.ISocketPack
	WebSocketPack websocket.IWebSocketPack
	IDiscover     discover.IDiscover
}

func NewManage(config *config.ManageConfig) IManage {
	return &Manage{
		IServer:       server.NewIServer(config.Server.Head),
		IMessage:      pack.NewIMessagePack(config),
		Db:            pack.NewIDbPack(config),
		IToolPack:     pack.NewIToolPack(),
		SocketPack:    socket.NewSocketPack(config),
		IDiscover:     discover.NewDiscover(),
		WebSocketPack: websocket.NewWebSocketPack(config),
	}
}

// 服务器管理
func (m *Manage) Server() iserver.IServer {
	return m.IServer
}

// 通信协议
func (m *Manage) Message() pack.IMessage {
	return m.IMessage
}

// 数据库
func (m *Manage) DB() pack.IDb {
	return m.Db
}

// 工具
func (m *Manage) Tool() pack.IToolPack {
	return m.IToolPack
}

// Socket
func (m *Manage) Socket() socket.ISocketPack {
	return m.SocketPack
}

// Socket
func (m *Manage) WebSocket() websocket.IWebSocketPack {
	return m.WebSocketPack
}

// 服务管理器
func (m *Manage) Discover() discover.IDiscover {
	return m.IDiscover
}
