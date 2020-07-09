package response

import (
	"KServer/server/manage"
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
	switch data.GetMsgId() {
	// 判断验证服务器是否判断成功 不成功则直接返回客户端

	case utils.OauthAccountSuccess:
		err := s.IManage.Client().GetClient(s.IManage.Message().Kafka().DataPack().GetClientConnId()).GetConn().
			SendMsg(data.GetId(), data.GetRawDate())
		if err != nil {
			fmt.Println("客户端回调消息失败")
		}

		// 升级客户端验证状态
		acc := &pd.Account{}
		s.IManage.Message().Kafka().DataPack().GetDate().ProtoBuf(acc)
		s.IManage.Client().UpgradeClient(s.IManage.Message().Kafka().DataPack().GetClientConnId(), acc)

	default:
		fmt.Println("我会执行吗")
		err := s.IManage.Client().GetClient(s.IManage.Message().Kafka().DataPack().GetClientConnId()).GetConn().
			SendMsg(data.GetId(), data.GetRawDate())
		if err != nil {
			fmt.Println("客户端回调消息失败")

		}
		acc := &pd.Account{}
		_ = data.GetDate().ProtoBuf(acc)
		s.IManage.Client().GetClientByUUID(acc.UUID).GetConn().Stop()
	}
}
