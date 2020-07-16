package server

import (
	"KServer/library/kiface/iserver"
	"fmt"
	"github.com/satori/go.uuid"
	"os"
	"os/signal"
	"syscall"
)

type Server struct {
	Id   string
	Host string
	Port string
}

func NewIServer(head string) iserver.IServer {
	return &Server{
		Id: head + uuid.NewV4().String(),
	}
}

func (s *Server) GetId() string {
	return s.Id
}
func (s *Server) SetAddr(host string, port string) {
	s.Host = host
	s.Port = port
}
func (s *Server) GetAddr() string {
	return s.Host + ":" + s.Port
}
func (s *Server) GetHost() string {
	return s.Host
}
func (s *Server) GetPort() string {
	return s.Port
}
func (s *Server) Start() {
	// 进入关闭监听
	fmt.Println(s.GetId(), "启动成功")

	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	signal.Notify(sigs, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()

	//fmt.Println("awaiting signal")

	<-done

	fmt.Println("Server Close...")
}
