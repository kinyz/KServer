package response

import (
	"KServer/manage"
	"KServer/server/utils"
	"fmt"
)

type ClientResponse struct {
	IManage manage.IManage
}

func NewClientResponse(m manage.IManage) *ClientResponse {
	return &ClientResponse{IManage: m}
}

// 用于接收客户端主题
func (c *ClientResponse) ResponseClient(data utils.IDataPack) {

	fmt.Println("客户端回调")
	//s.IManage.Message().DataPack().UnPack(req.GetData().Bytes())

	//fmt.Println("服务器全体收到消息", s.IManage.DataPack().GetMsgId())

	//switch s.IManage.DataPack().GetMsgId() {

}
