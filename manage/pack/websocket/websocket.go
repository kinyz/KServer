package websocket

import (
	"KServer/library/iface/iwebsocket"

	"KServer/library/websocket"
	"KServer/manage/config"
	"KServer/manage/pack/websocket/client"
)

type IWebSocketPack interface {
	Server() iwebsocket.IServer
	Client() client.IClientPack
}

type WebSocketPack struct {
	ServerPack iwebsocket.IServer
	ClientPack client.IClientPack
}

func NewWebSocketPack(conf *config.ManageConfig) IWebSocketPack {
	s := &WebSocketPack{}

	//fmt.Println(conf.Socket.Server)
	if conf.WebSocket.Server {
		s.ServerPack = websocket.NewWebsocket()
	}
	if conf.WebSocket.Client {
		s.ClientPack = client.NewIWsClientPack()
	}
	return s
}

func (s *WebSocketPack) Server() iwebsocket.IServer {
	return s.ServerPack
}
func (s *WebSocketPack) Client() client.IClientPack {
	return s.ClientPack
}
