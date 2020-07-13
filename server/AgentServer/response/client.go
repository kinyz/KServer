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

// 用于接收客户端主题
func (c *ClientResponse) ResponseClient(data proto.IDataPack) {

	fmt.Println("收到客户端信息", data.GetClientId())
	client := c.IManage.Socket().Client().GetClient(data.GetClientId())
	if client != nil {
		client.Send(data.GetData().Bytes())
		return
	}

	fmt.Println("客户端回调")

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
