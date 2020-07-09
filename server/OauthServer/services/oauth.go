package services

import (
	"KServer/server/manage"
	"KServer/server/utils"
	"KServer/server/utils/pd"
	"fmt"
)

type IOauth interface {
	ResponseOauth(data utils.IDataPack)
}

type Oauth struct {
	IManage manage.IManage
}

func NewOauth(i manage.IManage) IOauth {
	return &Oauth{IManage: i}
}

func (o *Oauth) ResponseOauth(data utils.IDataPack) {
	fmt.Println("收到网关信息", o.IManage.Message().Kafka().DataPack().GetMsgId())

	switch data.GetMsgId() {
	case utils.OauthAccount:
		acc := &pd.Account{}
		err := o.IManage.Message().Kafka().DataPack().GetDate().ProtoBuf(acc)
		if err != nil {
			fmt.Println("err=", err)
		}
		fmt.Println("数据接收", acc.UUID, acc.PassWord, acc.Token)
		o.IManage.Message().Kafka().Send().Async(data.GetServerId(), o.IManage.Server().GetId(),
			o.IManage.Message().Kafka().DataPack().Pack(utils.OauthMsgId, o.IManage.Message().Kafka().DataPack().GetClientId(), o.IManage.Message().Kafka().DataPack().GetServerId(),
				o.IManage.Message().Kafka().DataPack().GetClientConnId(), utils.OauthAccountSuccess, o.IManage.Message().Kafka().DataPack().GetDate().Bytes()))

	}

}
