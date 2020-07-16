package services

import (
	"KServer/manage"
	"KServer/proto"
	"KServer/server/utils/msg"
	"KServer/server/utils/pd"
	"fmt"
)

type IOauth interface {
	ResponseOauth(data proto.IDataPack)
}

type Oauth struct {
	IManage manage.IManage
}

func NewOauth(i manage.IManage) IOauth {
	return &Oauth{IManage: i}
}

func (o *Oauth) ResponseOauth(data proto.IDataPack) {
	fmt.Println("收到网关信息", data.GetId(), data.GetClientId(), data.GetServerId())
	switch data.GetMsgId() {
	case msg.OauthAccount:
		o.OauthAccount(data)
	case msg.OauthAccountClose:
		o.OauthAccountClose(data)
	}

}

func (o *Oauth) OauthAccount(data proto.IDataPack) {
	//	fmt.Println("步骤1")
	kafka := o.IManage.Message().Kafka()
	acc := &pd.Account{}
	data.GetData().ProtoBuf(acc)

	//fmt.Println("数据接收", acc.UUID, acc.PassWord, acc.Token)

	dbacc := &pd.Account{}
	err := o.IManage.DB().Redis().GetSlaveConn().Get(msg.ClientLoginInfoKey + data.GetClientId()).ProtoBuf(dbacc)
	//fmt.Println("步骤2",data.GetData().String())

	if err != nil {
		kafka.Send().Async(data.GetServerId(), o.IManage.Server().GetId(),
			o.IManage.Message().DataPack().Pack(
				data.GetId(),
				msg.OauthAccountSystemError,
				data.GetClientId(),
				o.IManage.Message().DataPack().GetServerId(),
				[]byte("系统错误")))
		return
	}
	//	fmt.Println("步骤3",acc.UUID,dbacc.UUID)

	if acc.UUID != dbacc.UUID {
		kafka.Send().Async(data.GetServerId(), o.IManage.Server().GetId(),
			o.IManage.Message().DataPack().Pack(
				data.GetId(),
				msg.OauthAccountNotFindError,
				data.GetClientId(),
				data.GetServerId(),
				[]byte("找不到账号")))
		return
	}
	//fmt.Println("步骤4")

	if acc.Token != dbacc.Token {
		kafka.Send().Async(data.GetServerId(), o.IManage.Server().GetId(),
			o.IManage.Message().DataPack().Pack(
				data.GetId(),
				msg.OauthAccountTokenError,
				data.GetClientId(),
				o.IManage.Server().GetId(),
				[]byte("Token已失效")))
		return
	}
	//	fmt.Println("步骤5", dbacc.Online)

	if dbacc.Online == msg.ClientOnline {
		kafka.Send().Async(data.GetServerId(), o.IManage.Server().GetId(),
			o.IManage.Message().DataPack().Pack(
				data.GetId(),
				msg.OauthAccountOnlineError,
				data.GetClientId(),
				o.IManage.Server().GetId(),
				[]byte("当前账号已在线")))
		return
	}
	//fmt.Println("步骤6")

	if dbacc.State != 0 {
		kafka.Send().Async(data.GetServerId(), o.IManage.Server().GetId(),
			o.IManage.Message().DataPack().Pack(
				data.GetId(),
				msg.OauthAccountAccountStateError,
				data.GetClientId(),
				o.IManage.Server().GetId(),
				[]byte("账号已被封停")))
		return
	}
	dbacc.Online = msg.OauthAccountOnline
	o.IManage.DB().Redis().GetMasterConn().Set(msg.ClientLoginInfoKey + acc.UUID).ProtoBuf(dbacc)
	kafka.Send().Async(data.GetServerId(), o.IManage.Server().GetId(),
		o.IManage.Message().DataPack().Pack(
			data.GetId(),
			msg.OauthAccountSuccess,
			data.GetClientId(),
			data.GetServerId(),
			[]byte("登陆成功")))
	//fmt.Println("步骤7")

}

func (o *Oauth) OauthAccountClose(data proto.IDataPack) {
	//fmt.Println("收到请求关闭", data.GetClientId())
	dbuser := &pd.Account{}

	err := o.IManage.DB().Redis().GetSlaveConn().Get(msg.ClientLoginInfoKey + data.GetClientId()).ProtoBuf(dbuser)

	//fmt.Println(dbuser)
	if err != nil {
		fmt.Println(err)
	}
	dbuser.Online = 0
	_, err = o.IManage.DB().Redis().GetMasterConn().Set(msg.ClientLoginInfoKey + data.GetClientId()).ProtoBuf(dbuser)

	if err != nil {
		fmt.Println("2", err)
	}
}
