package main

import (
	"KServer/library/iface/iwebsocket"
	"KServer/library/websocket"
	"fmt"
)

func main() {

	s := websocket.NewWebsocket()

	s.SetOnConnStart(action)
	s.SetOnConnStop(stop)
	s.AddHandle(100, Newhand())
	s.Serve()

	select {}
}

func action(conn iwebsocket.IConnection) {

	fmt.Println(conn.GetConnID(), "连接")

}
func stop(conn iwebsocket.IConnection) {
	fmt.Println(conn.GetConnID(), "离开")

}

func Newhand() *handelo {
	return &handelo{}
}

type handelo struct {
}

func (h *handelo) PreHandle(request iwebsocket.IRequest) {
	fmt.Println("PreHandle")

}

func (h *handelo) PostHandle(request iwebsocket.IRequest) {
	fmt.Println("PostHandle")
}
