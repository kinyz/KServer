package response

import (
	"KServer/manage"
	"KServer/server/utils"
	"fmt"
)

type AllServerResponse struct {
	IManage manage.IManage
}

func NewAllServerResponse(m manage.IManage) *AllServerResponse {
	return &AllServerResponse{IManage: m}
}

func (s *AllServerResponse) ResponseAllServer(data utils.IDataPack) {
	//s.IManage.Message().DataPack().UnPack(req.GetData().Bytes())

	fmt.Println("服务器全体收到消息", s.IManage.Message().Kafka().DataPack().GetDate().String())

	//switch s.IManage.DataPack().GetMsgId() {

}

// 移除客户端
func (s *AllServerResponse) RemoveClient(data utils.IDataPack) {
	c := s.IManage.Socket().Client().GetClient(data.GetClientId())
	if c == nil {
		fmt.Println("客户端不在此服务器")
		return
	}

	err := c.Send(data.GetId(), data.GetMsgId(), data.GetDate().Bytes())
	if err != nil {
		fmt.Println("客户端回调消息失败")
	}
	c.Stop()

}
