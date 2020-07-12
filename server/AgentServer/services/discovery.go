package services

import (
	"KServer/library/iface/isocket"
	"KServer/manage"
	"KServer/manage/discover/pd"
	"KServer/server/utils"
	"fmt"
)

type Discovery struct {
	IManage manage.IManage
}

func NewCustomHandle(m manage.IManage) *Discovery {
	return &Discovery{IManage: m}
}

func (c *Discovery) PreHandle(request isocket.IRequest) {
	if !c.IManage.Discover().CheckService(request.GetID()) {
		// 判断id是否存在
		request.GetConnection().Stop()
		return
	}

	data := c.IManage.Message().Kafka().DataPack().Pack(request.GetID(), request.GetMsgID(),
		c.IManage.Socket().Client().GetIdByConnId(request.GetConnection().GetConnID()),
		c.IManage.Server().GetId(), request.GetData())

	_, _, err := c.IManage.Message().Kafka().Send().Sync(
		c.IManage.Discover().GetTopic(request.GetID()),
		c.IManage.Server().GetId(), data)
	if err != nil {
		fmt.Println(request.GetID(), request.GetMsgID(), "转发失败")
	}
	fmt.Println("CustomHandle")
}

func (c *Discovery) PostHandle(request isocket.IRequest) {

	//fmt.Println("CustomHandle2")
}

// 用于服务中心注册服务
func (c *Discovery) DiscoverHandle(data utils.IDataPack) {
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
func (c *Discovery) ResponseAddService(data utils.IDataPack) {

	//fmt.Println("服务发现添加服务")
	d := &pd.Discovery{}
	err := data.GetDate().ProtoBuf(d)
	if err != nil {
		fmt.Println("服务发现解析失败")
		return
	}
	c.IManage.Discover().AddService(d.Id, d)
	fmt.Println(d.Id, d.Topic, "服务发现添加服务完成")
	fmt.Println(c.IManage.Discover().GetAllTopic())

}

// 用于服务中心删除服务
func (c *Discovery) ResponseDelService(data utils.IDataPack) {
	d := &pd.Discovery{}
	err := data.GetDate().ProtoBuf(d)
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
