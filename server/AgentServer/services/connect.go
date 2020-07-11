package services

import (
	"KServer/library/iface/ikafka"
	"KServer/library/iface/isocket"
	"KServer/library/socket"
	"KServer/manage"
	"KServer/manage/config"
	"KServer/server/AgentServer/response"
	"KServer/server/utils"
	"KServer/server/utils/pd"
	"fmt"
)

type Connect struct {
	socket.Handle
	IManage     manage.IManage
	KafkaConfig ikafka.IKafkaConf
}

func NewConnect(m manage.IManage) *Connect {
	return &Connect{IManage: m,
		KafkaConfig: config.NewKafkaConfig(utils.KafkaConFile)}
}

func (c *Connect) PreHandle(request isocket.IRequest) {

	//fmt.Println(request.GetID())

	switch request.GetMsgID() {
	case utils.OauthAccount:
		_ = c.IManage.Message().Kafka().DataPack().UnPack(request.GetData())
		acc := &pd.Account{}
		_ = c.IManage.Message().Kafka().DataPack().GetDate().ProtoBuf(acc)

		//fmt.Println("当前客户端状态", c.IManage.Socket().Client().GetState(acc.UUID))
		if c.IManage.Socket().Client().GetState(acc.UUID) {
			request.GetConnection().SendMsg(request.GetID(), utils.ClientOnlineError, []byte("当前账号已在线"))
			request.GetConnection().Stop()
			if c.IManage.Socket().Client().GetState(acc.UUID) {
				c.IManage.Socket().Client().GetClient(acc.UUID).Send(request.GetID(), utils.ClientConnectConnIdError, []byte("当前账号已在其他地方登陆"))
				c.IManage.Socket().Client().GetClient(acc.UUID).Stop()
			}
			return
		}
		// 加入客户端管理器
		c.IManage.Socket().Client().AddClient(acc.UUID, acc.Token, request.GetConnection())

		// kafka
		clientResponse := response.NewClientResponse(c.IManage)

		clientTopic := utils.ClientTopic + acc.UUID // 客户端消费头
		clientListenTopic := []string{
			clientTopic,
		}
		if !c.IManage.Socket().Client().GetState(acc.UUID) {
			return
		}
		// 新增接收客户端的kafka路由
		c.IManage.Message().Kafka().AddRouter(clientTopic, utils.ClientNotifyId, clientResponse.ResponseClient)
		c.IManage.Message().Kafka().AddRouter(clientTopic, utils.ClientRemove, clientResponse.ResponseRemoveClient)

		// 启动客户端所需启动的监听
		go func() {
			c.IManage.Socket().Client().SetClose(acc.UUID, c.IManage.Message().Kafka().StartOtherListen(
				clientListenTopic,
				[]string{c.KafkaConfig.GetAddr()},
				c.IManage.Tool().Encrypt().NewUuid(),
				-1))

			data := c.IManage.Message().Kafka().DataPack().Pack(request.GetID(),
				request.GetMsgID(),
				acc.UUID,
				c.IManage.Server().GetId(),
				c.IManage.Message().Kafka().DataPack().GetDate().Bytes())
			//c.IManage.Message().DataPack()
			c.IManage.Message().Kafka().Send().Async(utils.OauthTopic, c.IManage.Server().GetId(), data)
		}()

	}

}

func (c *Connect) PostHandle(request isocket.IRequest) {

}

//创建连接的时候执行
func (c *Connect) DoConnectionBegin(conn isocket.IConnection) {

	//zlog.Debug("[创建连接] IP:", conn.RemoteAddr(), " ConnId:", conn.GetConnID())
	//conn.SetProperty(GlobalMessage.ClientConnectOauthKey, false)
	//c.IManage.Client().GetClientByConnId(conn.GetConnID()).SetConn(conn)
	//c.IManage.Client().GetClient(conn.GetConnID()).SetUUID("")
	err := conn.SendMsg(utils.OauthId, utils.OauthAccount, []byte("DoConnection BEGIN..."))
	if err != nil {
		//	zlog.Error(err)
	}

}

//连接断开的时候执行
func (c *Connect) DoConnectionLost(conn isocket.IConnection) {

	//c.IManage.Client().RemoveClient(conn.GetConnID())
	fmt.Println("[断开连接] IP:", conn.RemoteAddr(), " ConnId:", conn.GetConnID(), " UUID:", c.IManage.Socket().Client().GetIdByConnId(conn.GetConnID()))

	c.IManage.Message().Kafka().Send().Sync(utils.OauthTopic, c.IManage.Server().GetId(),
		c.IManage.Message().Kafka().DataPack().Pack(
			utils.OauthId,
			utils.OauthAccountClose,
			c.IManage.Socket().Client().GetIdByConnId(conn.GetConnID()),
			c.IManage.Server().GetId(),
			[]byte("close")))

	//c.IManage.Message().DataPack()
	//	c.IManage.Message().Kafka().Send().Async(utils.OauthTopic, c.IManage.Server().GetId(), data)

	c.IManage.Socket().Client().Remove(conn.GetConnID())

}

func (c *Connect) ResponseOauth(data utils.IDataPack) {

	client := c.IManage.Socket().Client().GetClient(data.GetClientId())
	if client == nil {
		//fmt.Println("客户端不存在")
		return
	}
	//fmt.Println("验证回调",data.GetMsgId())
	switch data.GetMsgId() {
	// 判断验证服务器是否判断成功 不成功则直接返回客户端

	case utils.OauthAccountSuccess:

		err := client.Send(data.GetId(), data.GetMsgId(), data.GetDate().Bytes())
		if err != nil {
			fmt.Println("客户端回调消息失败")
		}

	default:
		//	fmt.Println("我会执行吗")
		err := client.Send(data.GetId(), data.GetMsgId(), data.GetDate().Bytes())
		if err != nil {
			fmt.Println("客户端回调消息失败")
		}
		client.Stop()

	}

}
