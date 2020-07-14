package services

import (
	"KServer/library/iface/isocket"
	"KServer/manage"
	"KServer/proto"
	"KServer/server/utils/msg"
	"fmt"
)

type SocketDiscovery struct {
	IManage manage.IManage
}

func NewSocketDiscovery(m manage.IManage) *SocketDiscovery {
	return &SocketDiscovery{IManage: m}
}

func (c *SocketDiscovery) PreHandle(request isocket.IRequest) {
	if !c.IManage.Discover().CheckService(request.GetMessage().GetId()) {
		// 判断id是否存在
		request.GetConnection().SendBuffMsg([]byte("无服务"))
		return
	}
	data := c.IManage.Message().DataPack().Pack(request.GetMessage().GetId(), request.GetMessage().GetMsgId(),
		c.IManage.Socket().Client().GetIdByConnId(request.GetConnection().GetConnID()),
		c.IManage.Server().GetId(), request.GetMessage().GetData())

	_, _, err := c.IManage.Message().Kafka().Send().Sync(
		c.IManage.Discover().GetTopic(request.GetMessage().GetId()),
		c.IManage.Server().GetId(), data)
	if err != nil {
		fmt.Println(request.GetMessage().GetId(), request.GetMessage().GetMsgId(), "转发失败")
	}
	//fmt.Println("CustomHandle")
}

func (c *SocketDiscovery) PostHandle(request isocket.IRequest) {

	//fmt.Println("CustomHandle2")
}

// 用于服务中心头
func (c *SocketDiscovery) DiscoverHandle(data proto.IDataPack) {

	fmt.Println("收到服务变化", data.GetMsgId())
	switch data.GetMsgId() {
	case msg.ServiceDiscoveryRegister:
		{
			c.ResponseAddService(data)
		}
	case msg.ServiceDiscoveryCloseService:
		c.ResponseDelService(data)
	}
}

// 用于服务中心注册服务
func (c *SocketDiscovery) ResponseAddService(data proto.IDataPack) {

	//fmt.Println("服务发现添加服务")
	d := &proto.Discovery{}
	err := data.GetData().ProtoBuf(d)
	if err != nil {
		fmt.Println("服务发现解析失败")
		return
	}
	c.IManage.Discover().AddService(d.Id, d)
	fmt.Println(d.Id, d.Topic, "服务发现添加服务完成")
	//fmt.Println(c.IManage.Discover().GetAllTopic())

}

// 用于服务中心删除服务
func (c *SocketDiscovery) ResponseDelService(data proto.IDataPack) {
	d := &proto.Discovery{}
	err := data.GetData().ProtoBuf(d)
	//fmt.Println("服务发现删除服务")
	if err != nil {
		fmt.Println("服务发现解析失败")
		return
	}
	if c.IManage.Discover().CheckService(d.Id) {
		c.IManage.Discover().DelService(d.Id)
		fmt.Println(d.Id, d.Topic, "服务发现删除服务完成")
		//fmt.Println(c.IManage.Discover().GetAllTopic())

	}

}
