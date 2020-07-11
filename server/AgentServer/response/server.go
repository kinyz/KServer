package response

import (
	"KServer/manage"
	"KServer/server/utils"
	"KServer/server/utils/pd"
	"fmt"
)

type ServerResponse struct {
	IManage manage.IManage
}

func NewIServerResponse(m manage.IManage) *ServerResponse {
	return &ServerResponse{IManage: m}
}
func (s *ServerResponse) ResponseOauth(data utils.IDataPack) {
	c := s.IManage.Socket().Client().GetClient(data.GetClientId())
	if c == nil {
		return
	}
	switch data.GetMsgId() {
	// 判断验证服务器是否判断成功 不成功则直接返回客户端

	case utils.OauthAccountSuccess:
		err := c.Send(data.GetId(), data.GetRawDate())
		if err != nil {
			fmt.Println("客户端回调消息失败")
		}

	default:
		fmt.Println("我会执行吗")
		err := c.Send(data.GetId(), data.GetRawDate())
		if err != nil {
			fmt.Println("客户端回调消息失败")

		}
		acc := &pd.Account{}
		_ = data.GetDate().ProtoBuf(acc)
		c.Stop()
	}
}
