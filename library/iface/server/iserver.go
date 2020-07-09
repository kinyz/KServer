package server

type IServer interface {
	GetId() string
	SetAddr(host string, port string)
	GetAddr() string
	GetHost() string
	GetPort() string
	Start()
}
