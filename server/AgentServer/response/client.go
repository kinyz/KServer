package response

import (
	"KServer/manage"
	"KServer/proto"
	"fmt"
)

type ClientResponse struct {
	IManage manage.IManage
}

func NewClientResponse(m manage.IManage) *ClientResponse {
	return &ClientResponse{IManage: m}
}

// 用于接收客户端的自定义头
func (c *ClientResponse) ResponseClient(data proto.IDataPack) {
	fmt.Println("收到来自客户端的回调1", data.GetClientId())

	if c.IManage.Socket().Client().GetState(data.GetClientId()) {
		c.IManage.Socket().Client().GetClient(data.GetClientId()).Send(data.GetRawData())
		return
	}
	fmt.Println("收到来自客户端的回调2", data.GetData().String())

}

// 用于接收客户端主题
func (c *ClientResponse) ResponseRemoveClient(data proto.IDataPack) {

	client := c.IManage.Socket().Client().GetClient(data.GetClientId())
	if client != nil {
		client.SendBuff(data.GetData().Bytes())
		client.Stop()
		return
	}

	fmt.Println("客户端回调")

}
