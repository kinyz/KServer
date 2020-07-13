package services

import (
	"KServer/manage"
	"KServer/proto"
	"KServer/server/utils"
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
	fmt.Println("收到网关信息", o.IManage.Message().Kafka().DataPack().GetMsgId())

	switch data.GetMsgId() {
	case utils.OauthAccount:
		kafka := o.IManage.Message().Kafka()
		acc := &pd.Account{}
		err := kafka.DataPack().GetData().ProtoBuf(acc)
		if err != nil {
			fmt.Println("err=", err)
			return
		}
		//fmt.Println("数据接收", acc.UUID, acc.PassWord, acc.Token)

		dbacc := &pd.Account{}
		err = o.IManage.DB().Redis().GetSlaveConn().Get(utils.ClientLoginInfoKey + acc.UUID).ProtoBuf(dbacc)

		if err != nil {
			kafka.Send().Async(data.GetServerId(), o.IManage.Server().GetId(),
				kafka.DataPack().Pack(
					data.GetId(),
					utils.OauthAccountSystemError,
					acc.UUID,
					o.IManage.Message().Kafka().DataPack().GetServerId(),
					[]byte("系统错误")))
			return
		}
		if acc.UUID != dbacc.UUID {
			kafka.Send().Async(data.GetServerId(), o.IManage.Server().GetId(),
				kafka.DataPack().Pack(
					data.GetId(),
					utils.OauthAccountNotFindError,
					acc.UUID,
					o.IManage.Message().Kafka().DataPack().GetServerId(),
					[]byte("找不到账号")))
			return
		}
		if acc.Token != dbacc.Token {
			kafka.Send().Async(data.GetServerId(), o.IManage.Server().GetId(),
				kafka.DataPack().Pack(
					data.GetId(),
					utils.OauthAccountTokenError,
					acc.UUID,
					o.IManage.Message().Kafka().DataPack().GetServerId(),
					[]byte("Token已失效")))
			return
		}
		if dbacc.Online == utils.ClientOnline {
			kafka.Send().Async(data.GetServerId(), o.IManage.Server().GetId(),
				kafka.DataPack().Pack(
					data.GetId(),
					utils.OauthAccountOnlineError,
					acc.UUID,
					o.IManage.Message().Kafka().DataPack().GetServerId(),
					[]byte("当前账号已在线")))
			return
		}

		if dbacc.State != 0 {
			kafka.Send().Async(data.GetServerId(), o.IManage.Server().GetId(),
				kafka.DataPack().Pack(
					data.GetId(),
					utils.OauthAccountAccountStateError,
					acc.UUID,
					o.IManage.Message().Kafka().DataPack().GetServerId(),
					[]byte("账号已被封停")))
			return
		}
		dbacc.Online = utils.OauthAccountOnline
		o.IManage.DB().Redis().GetMasterConn().Set(utils.ClientLoginInfoKey + acc.UUID).ProtoBuf(dbacc)
		kafka.Send().Async(data.GetServerId(), o.IManage.Server().GetId(),
			kafka.DataPack().Pack(
				data.GetId(),
				utils.OauthAccountSuccess,
				acc.UUID,
				o.IManage.Message().Kafka().DataPack().GetServerId(),
				o.IManage.Message().Kafka().DataPack().GetData().Bytes()))
	case utils.OauthAccountClose:

		fmt.Println("收到请求关闭", data.GetClientId())
		dbuser := &pd.Account{}

		err := o.IManage.DB().Redis().GetSlaveConn().Get(utils.ClientLoginInfoKey + data.GetClientId()).ProtoBuf(dbuser)

		//fmt.Println(dbuser)
		if err != nil {
			fmt.Println(err)
		}
		dbuser.Online = 0
		_, err = o.IManage.DB().Redis().GetMasterConn().Set(utils.ClientLoginInfoKey + data.GetClientId()).ProtoBuf(dbuser)

		if err != nil {
			fmt.Println("2", err)
		}

	}

}
