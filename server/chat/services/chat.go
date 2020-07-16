package services

import (
	"KServer/manage"
	"KServer/proto"
	"KServer/server/utils/msg"
	"KServer/server/utils/pd"
	"fmt"
)

type Chat struct {
	manage.IManage
	msg pd.Chat
}

func NewChat(m manage.IManage) *Chat {
	return &Chat{IManage: m}
}

func (c *Chat) ResponseChat(data proto.IDataPack) {
	switch data.GetMsgId() {
	// 发送世界
	case msg.ChatWord:
		{
			c.SendWord(data)
		}
	// 发送私人
	case msg.ChatPrivate:
		{
			c.SendPrivate(data)
		}

	// 发送组
	case msg.ChatGroup:
		{
			c.SendGroup(data)
		}
	}
}

func (c *Chat) SendWord(data proto.IDataPack) {
	//c.IManage.Message().DataPack().GetData().ProtoBuf(c.msg)
	fmt.Println("发送全体")

	push := c.Message().DataPack().Pack(msg.AgentAllServerId, msg.AgentSendAllClient, data.GetClientId(), c.Server().GetId(), data.GetData().Bytes())
	c.Message().Kafka().Send().Async(msg.AgentServerAllTopic, c.Server().GetId(), push)
}
func (c *Chat) SendPrivate(data proto.IDataPack) {
	fmt.Println("发送私人1")
	err := data.GetData().ProtoBuf(&c.msg)
	fmt.Println("发送私人2", &c.msg)
	if err != nil {
		// 写错误日志
		fmt.Println("发送私人3", err)

		return
	}
	_, _, err = c.Message().Kafka().Send().Sync(msg.ClientTopic+c.msg.Send, c.Server().GetId(), data.GetRawData())
	if err != nil {
		// 写错误日志
		return
	}
}
func (c *Chat) SendGroup(data proto.IDataPack) {
	err := data.GetData().ProtoBuf(&c.msg)
	if err != nil {
		// 写错误日志
		return
	}
	_, _, err = c.Message().Kafka().Send().Sync(msg.GroupTopic+c.msg.Send, c.Server().GetId(), data.GetRawData())
	if err != nil {
		// 写错误日志
		return
	}

}
