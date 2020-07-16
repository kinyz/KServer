package socket

import (
	"KServer/library/kiface/isocket"
	"KServer/library/socket"
	"KServer/manage/config"
	"KServer/manage/pack/socket/client"
	"fmt"
)

type ISocketPack interface {
	Server() isocket.IServer
	Client() client.IClientPack
}

type SocketPack struct {
	ServerPack isocket.IServer
	ClientPack client.IClientPack
}

func NewSocketPack(conf *config.ManageConfig) ISocketPack {
	s := &SocketPack{}

	if conf.Socket.Server {
		fmt.Println(conf.Socket.Server)

		s.ServerPack = socket.NewSocket()
	}
	if conf.Socket.Client {
		s.ClientPack = client.NewIClientPack()
	}
	return s
}

func (s *SocketPack) Server() isocket.IServer {
	return s.ServerPack
}
func (s *SocketPack) Client() client.IClientPack {
	return s.ClientPack
}
