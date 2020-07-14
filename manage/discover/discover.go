package discover

import (
	"KServer/proto"
	"fmt"
)

type IDiscover interface {
	/*
		获取服务Topic
		id uint32 服务id
		返回值 string 服务Topic
	*/
	GetTopic(id uint32) string
	/*
		获取服务Host
		id uint32 服务id
		返回值 string 服务Host
	*/
	GetHost(id uint32) string
	/*
		获取服务Port
		id uint32 服务id
		返回值 string 服务Prot
	*/
	GetPort(id uint32) string
	/*
		获取服务Type
		id uint32 服务id
		返回值 string 服务Type
	*/
	GetType(id uint32) string
	/*
		增加服务
		id uint32 服务id
		service *pd.Discovery 服务内容
	*/
	AddService(id uint32, service *proto.Discovery)
	/*
		删除服务
		id uint32 服务id
	*/
	DelService(id uint32)
	/*
		检查服务
		id uint32 服务id
		返回值 bool 服务健康状态
	*/
	CheckService(id uint32) bool
	/*
		获取所有服务
		返回值 *pd.Discovery
	*/
	GetAllTopic() map[uint32]*proto.Discovery
}
type Discover struct {
	Topic map[uint32]*proto.Discovery
}

func NewDiscover() IDiscover {
	return &Discover{
		Topic: make(map[uint32]*proto.Discovery),

		//Id: make(map[string]uint32),
	}
}

func (d *Discover) GetTopic(id uint32) string {
	return d.Topic[id].Topic
}

func (d *Discover) GetHost(id uint32) string {
	return d.Topic[id].Host
}

func (d *Discover) GetPort(id uint32) string {
	return d.Topic[id].Port
}

func (d *Discover) GetType(id uint32) string {
	return d.Topic[id].Type
}

func (d *Discover) AddService(id uint32, service *proto.Discovery) {
	//if d.Topic[id] != nil {
	//return
	//}
	d.Topic[id] = service
	fmt.Println("添加服务", id, d.Topic[id])
}
func (d *Discover) DelService(id uint32) {
	if d.Topic[id] != nil {
		fmt.Println("删除服务", id, d.Topic[id])
		delete(d.Topic, id)
	}

}

func (d *Discover) CheckService(id uint32) bool {
	if d.Topic[id] != nil {
		return true
	}
	return false
}

func (d *Discover) GetAllTopic() map[uint32]*proto.Discovery {
	return d.Topic
}
