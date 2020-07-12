package socket

import (
	"KServer/library/iface/isocket"
	"fmt"
)

type Handle struct {
	Id uint32
	//Msg     map[uint32]func(response ziface.IResponse)
	Handle map[uint32]isocket.IHandle //存放每个Id 所对应的处理方法的map属性
	//Response ziface.IResponse

}

func NewIAgreement(id uint32, handle isocket.IHandle) *Handle {
	a := &Handle{
		Id: id,
		//Response: &Response{},
		//	Msg:     make(map[uint32]func(ziface.IResponse)),
		Handle: make(map[uint32]isocket.IHandle),
		//Response: make(map[uint32]ziface.IResponse),
	}
	a.Handle[id] = handle
	//a.Id[id] = handle
	return a
}

func (a *Handle) PreHandle(req isocket.IRequest) {

	handle, ok := a.Handle[req.GetID()]
	if ok {
		handle.PreHandle(req)
		return
	}

	fmt.Println("Agreement PreHandle Not Fount")

}

func (a *Handle) PostHandle(req isocket.IRequest) {
	handle, ok := a.Handle[req.GetID()]
	if ok {
		handle.PostHandle(req)
		return
	}

	fmt.Println("Agreement PostHandle Not Fount")

}
