package response

import (
	"KServer/server/manage"
	"KServer/server/utils"
)

type AllServerResponse struct {
	IManage manage.IManage
}

func NewAllServerResponse(m manage.IManage) *AllServerResponse {
	return &AllServerResponse{IManage: m}
}

func (s *AllServerResponse) ResponseAllStop(data utils.IDataPack) {
	//s.IManage.Message().DataPack().UnPack(req.GetData().Bytes())

	//fmt.Println("服务器全体收到消息", s.IManage.DataPack().GetMsgId())

	//switch s.IManage.DataPack().GetMsgId() {

}
