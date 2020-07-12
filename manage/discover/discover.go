package discover

import (
	"KServer/manage/discover/pd"
	"fmt"
)

type IDiscover interface {
	GetTopic(id uint32) string
	GetHost(id uint32) string
	GetPort(id uint32) string
	GetType(id uint32) string
	AddService(id uint32, service *pd.Discovery)
	DelService(id uint32)

	CheckService(id uint32) bool
	GetAllTopic() map[uint32]*pd.Discovery
}
type Discover struct {
	Topic map[uint32]*pd.Discovery
}

func NewDiscover() IDiscover {
	return &Discover{
		Topic: make(map[uint32]*pd.Discovery),

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

func (d *Discover) AddService(id uint32, service *pd.Discovery) {
	if d.Topic[id] != nil {
		return
	}
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

func (d *Discover) GetAllTopic() map[uint32]*pd.Discovery {
	return d.Topic
}
