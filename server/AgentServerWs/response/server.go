package response

import (
	"KServer/manage"
	"KServer/proto"
)

type ServerResponse struct {
	IManage manage.IManage
}

func NewIServerResponse(m manage.IManage) *ServerResponse {
	return &ServerResponse{IManage: m}
}

func (s *ServerResponse) SendAllClient(data proto.IDataPack) {

	s.IManage.WebSocket().Client().SendAll(data.GetData().Bytes())
}
