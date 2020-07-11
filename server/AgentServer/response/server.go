package response

import (
	"KServer/manage"
	"KServer/server/utils"
)

type ServerResponse struct {
	IManage manage.IManage
}

func NewIServerResponse(m manage.IManage) *ServerResponse {
	return &ServerResponse{IManage: m}
}

func (s *ServerResponse) SendAllClient(data utils.IDataPack) {

	s.IManage.Socket().Client().SendAll(data.GetId(), data.GetMsgId(), data.GetDate().Bytes())
}
