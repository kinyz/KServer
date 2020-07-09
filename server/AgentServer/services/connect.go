package services

import (
	"KServer/library/iface/socket/ziface"
	"KServer/library/socket/zlog"
	"KServer/library/socket/znet"
	"KServer/server/manage"
	"KServer/server/utils"
	"KServer/server/utils/pd"
	"fmt"
)

type Connect struct {
	znet.BaseRouter
	IManage manage.IManage
}

func (c *Connect) Response(data utils.IDataPack) {
	panic("implement me")
}

func NewConnect(m manage.IManage) *Connect {
	return &Connect{IManage: m}
}

func (c *Connect) PreHandle(request ziface.IRequest) {

	c.IManage.Message().Kafka().DataPack().UnPack(request.GetData())
	acc := &pd.Account{}
	c.IManage.Message().Kafka().DataPack().GetDate().ProtoBuf(acc)
	//fmt.Println()
	switch c.IManage.Message().Kafka().DataPack().GetMsgId() {
	case utils.OauthAccount:
		data := c.IManage.Message().Kafka().DataPack().Pack(request.GetMsgID(), "",
			c.IManage.Server().GetId(), request.GetConnection().GetConnID(), c.IManage.Message().Kafka().DataPack().GetMsgId(),
			c.IManage.Message().Kafka().DataPack().GetDate().Bytes())
		//c.IManage.Message().DataPack()
		c.IManage.Message().Kafka().Send().Async(utils.OauthTopic, c.IManage.Server().GetId(), data)
	}

}

//创建连接的时候执行
func (c *Connect) DoConnectionBegin(conn ziface.IConnection) {

	zlog.Debug("[创建连接] IP:", conn.RemoteAddr(), " ConnId:", conn.GetConnID())
	//conn.SetProperty(GlobalMessage.ClientConnectOauthKey, false)
	if !c.IManage.Client().AddClient(conn.GetConnID()) {
		_ = conn.SendMsg(utils.ClientConnectOauthError, []byte("登陆失败"))
		conn.Stop()
		return
	}
	c.IManage.Client().GetClient(conn.GetConnID()).SetConn(conn)
	err := conn.SendMsg(utils.OauthMsgId, []byte("DoConnection BEGIN..."))
	if err != nil {
		zlog.Error(err)
	}

}

//连接断开的时候执行
func (c *Connect) DoConnectionLost(conn ziface.IConnection) {

	c.IManage.Client().RemoveClient(conn.GetConnID())
	zlog.Debug("[断开连接] IP:", conn.RemoteAddr(), " ConnId:", conn.GetConnID())

}

func (c *Connect) ResponseOauth(data utils.IDataPack) {
	switch data.GetMsgId() {
	// 判断验证服务器是否判断成功 不成功则直接返回客户端

	case utils.OauthAccountSuccess:
		err := c.IManage.Client().GetClient(data.GetClientConnId()).GetConn().
			SendMsg(data.GetId(), data.GetRawDate())
		if err != nil {
			fmt.Println("客户端回调消息失败")
		}

		// 升级客户端验证状态
		acc := &pd.Account{}
		data.GetDate().ProtoBuf(acc)
		c.IManage.Client().UpgradeClient(data.GetClientConnId(), acc)

	default:
		fmt.Println("我会执行吗")
		err := c.IManage.Client().GetClient(data.GetClientConnId()).GetConn().
			SendMsg(data.GetId(), data.GetRawDate())
		if err != nil {
			fmt.Println("客户端回调消息失败")

		}
		acc := &pd.Account{}
		_ = data.GetDate().ProtoBuf(acc)
		c.IManage.Client().GetClientByUUID(acc.UUID).GetConn().Stop()

	}

}
