package services

import (
	"KServer/library/kiface/ikafka"
	"KServer/library/kiface/iwebsocket"
	"KServer/library/websocket"
	"KServer/manage"
	"KServer/manage/config"
	"KServer/proto"
	"KServer/server/utils"
	"KServer/server/utils/msg"
	"KServer/server/utils/pd"
	"KServer/server/websocket/response"
	"fmt"
)

type Connect struct {
	websocket.Handle
	IManage     manage.IManage
	KafkaConfig ikafka.IKafkaConf
}

func NewConnect(m manage.IManage) *Connect {
	return &Connect{IManage: m,
		KafkaConfig: config.NewKafkaConfig(utils.KafkaConFile)}
}

func (c *Connect) PreHandle(request iwebsocket.IRequest) {

	//fmt.Println(request.GetId())

	switch request.GetMessage().GetMsgId() {
	case msg.OauthAccount:
		//_ = c.IManage.Message().Kafka().DataPack().UnPack(request.GetMessage().GetData())
		acc := &pd.Account{}
		c.IManage.Tool().Protobuf().Decode(request.GetMessage().GetData(), acc)

		//fmt.Println("当前客户端状态", c.IManage.Socket().Client().GetState(acc.UUID))
		if c.IManage.WebSocket().Client().GetState(acc.UUID) {
			request.GetConnection().SendMsg([]byte("当前账号已在线"))
			request.GetConnection().Stop()
			if c.IManage.WebSocket().Client().GetState(acc.UUID) {
				c.IManage.WebSocket().Client().GetClient(acc.UUID).Send(request.GetMessage().GetRawData())
				c.IManage.WebSocket().Client().GetClient(acc.UUID).Stop()
			}
			return
		}
		// 加入客户端管理器
		c.IManage.WebSocket().Client().AddClient(acc.UUID, acc.Token, request.GetConnection())

		// pack一个向验证服务器请求验证的包
		data := c.IManage.Message().DataPack().Pack(
			request.GetMessage().GetId(),
			request.GetMessage().GetMsgId(),
			acc.UUID,
			c.IManage.Server().GetId(),
			request.GetMessage().GetData())

		//	fmt.Println(string(request.GetMessage().GetData()))
		c.IManage.Message().Kafka().Send().Async(msg.OauthTopic, c.IManage.Server().GetId(), data)

	}

}

func (c *Connect) PostHandle(request iwebsocket.IRequest) {

}

//创建连接的时候执行
func (c *Connect) DoConnectionBegin(conn iwebsocket.IConnection) {

	//zlog.Debug("[创建连接] IP:", conn.RemoteAddr(), " ConnId:", conn.GetConnID())
	//conn.SetProperty(GlobalMessage.ClientConnectOauthKey, false)
	//c.IManage.Client().GetClientByConnId(conn.GetConnID()).SetConn(conn)
	//c.IManage.Client().GetClient(conn.GetConnID()).SetUUID("")
	err := conn.SendMsg([]byte("DoConnection BEGIN..."))
	if err != nil {
		//	zlog.Error(err)
	}

}

//连接断开的时候执行
func (c *Connect) DoConnectionLost(conn iwebsocket.IConnection) {

	//c.IManage.Client().RemoveClient(conn.GetConnID())
	fmt.Println("[断开连接] IP:", conn.RemoteAddr(), " ConnId:", conn.GetConnID(), " UUID:", c.IManage.WebSocket().Client().GetIdByConnId(conn.GetConnID()))

	c.IManage.Message().Kafka().Send().Sync(msg.OauthTopic, c.IManage.Server().GetId(),
		c.IManage.Message().DataPack().Pack(
			msg.OauthId,
			msg.OauthAccountClose,
			c.IManage.WebSocket().Client().GetIdByConnId(conn.GetConnID()),
			c.IManage.Server().GetId(),
			[]byte("close")))

	//c.IManage.Message().DataPack()
	//	c.IManage.Message().Kafka().Send().Async(utils.OauthTopic, c.IManage.Server().GetId(), data)

	if c.IManage.Message().Kafka().RemoveCustomRouter(msg.ClientTopic + c.IManage.WebSocket().Client().GetIdByConnId(conn.GetConnID())) {
		fmt.Println("移除客户端路由：", c.IManage.WebSocket().Client().GetIdByConnId(conn.GetConnID()))
	}

	c.IManage.WebSocket().Client().Remove(conn.GetConnID())

}

func (c *Connect) ResponseOauth(data proto.IDataPack) {

	client := c.IManage.WebSocket().Client().GetClient(data.GetClientId())
	if client == nil {
		//fmt.Println("客户端不存在")
		return
	}
	//fmt.Println("验证回调",data.GetMsgId())
	switch data.GetMsgId() {
	// 判断验证服务器是否判断成功 不成功则直接返回客户端

	case msg.OauthAccountSuccess:

		// kafka回调验证成功
		clientResponse := response.NewClientResponse(c.IManage)

		clientTopic := msg.ClientTopic + data.GetClientId() // 客户端消费头
		clientListenTopic := []string{
			clientTopic,
		}
		if !c.IManage.WebSocket().Client().GetState(data.GetClientId()) {
			return
		}

		// 新增接收客户端的kafka路由
		c.IManage.Message().Kafka().AddCustomRouter(clientTopic, clientResponse.ResponseClient)
		//c.IManage.Message().Kafka().AddRouter(clientTopic, utils.ClientNotifyId, clientResponse.ResponseClient)
		//c.IManage.Message().Kafka().AddCustomRouter(clientTopic, clientResponse.ResponseRemoveClient)

		// 启动客户端所需启动的监听

		c.IManage.WebSocket().Client().SetClose(data.GetClientId(), c.IManage.Message().Kafka().StartCustomListen(
			clientListenTopic,
			[]string{c.KafkaConfig.GetAddr()},
			c.IManage.Tool().Encrypt().NewUuid(),
			-1))

		err := client.Send(data.GetRawData())
		if err != nil {
			fmt.Println("客户端回调消息失败")
			client.Stop()
		}

	default:
		//	fmt.Println("我会执行吗")
		err := client.Send(data.GetRawData())
		if err != nil {
			fmt.Println("客户端回调消息失败")
		}
		client.Stop()

	}

}
