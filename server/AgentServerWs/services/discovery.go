package services

import (
	"KServer/library/iface/iwebsocket"
	"KServer/manage"
	"KServer/manage/discover/pd"
	"KServer/proto"
	"KServer/server/utils"
	"fmt"
)

type WebSocketDiscovery struct {
	IManage manage.IManage
}

func NewWebSocketCustomHandle(m manage.IManage) *WebSocketDiscovery {
	return &WebSocketDiscovery{IManage: m}
}

func (c *WebSocketDiscovery) PreHandle(request iwebsocket.IRequest) {
	if !c.IManage.Discover().CheckService(request.GetMessage().GetId()) {
		// 判断id是否存在
		request.GetConnection().Stop()
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
	fmt.Println("CustomHandle")
}

func (c *WebSocketDiscovery) PostHandle(request iwebsocket.IRequest) {

	//fmt.Println("CustomHandle2")
}

// 用于服务中心注册服务
func (c *WebSocketDiscovery) DiscoverHandle(data proto.IDataPack) {
	switch data.GetMsgId() {
	case utils.ServiceDiscoveryRegister:
		{
			c.ResponseAddService(data)
		}
	case utils.ServiceDiscoveryCloseService:
		c.ResponseDelService(data)
	}
}

// 用于服务中心注册服务
func (c *WebSocketDiscovery) ResponseAddService(data proto.IDataPack) {

	//fmt.Println("服务发现添加服务")
	d := &pd.Discovery{}
	err := data.GetData().ProtoBuf(d)
	if err != nil {
		fmt.Println("服务发现解析失败")
		return
	}
	c.IManage.Discover().AddService(d.Id, d)
	fmt.Println(d.Id, d.Topic, "服务发现添加服务完成")
	fmt.Println(c.IManage.Discover().GetAllTopic())

}

// 用于服务中心删除服务
func (c *WebSocketDiscovery) ResponseDelService(data proto.IDataPack) {
	d := &pd.Discovery{}
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
