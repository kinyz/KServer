package services

import (
	"KServer/manage"
	"KServer/proto"
	"KServer/server/utils"
	"KServer/server/utils/msg"
	"fmt"
	"gopkg.in/mgo.v2/bson"
)

type Service struct {
	IManage manage.IManage
}

func NewServiceDiscovery(m manage.IManage) *Service {

	return &Service{IManage: m}
}

// 服务头
func (s *Service) ServiceHandle(data proto.IDataPack) {
	fmt.Println("接收消息", data.GetMsgId())
	switch data.GetMsgId() {
	case msg.ServiceDiscoveryRegister:
		{
			s.RegisterService(data)
		}

	case msg.ServiceDiscoveryLogoutService:
		{
			s.LogoutService(data)
		}
	case msg.ServiceDiscoveryCheckAllService:
		{
			s.CheckAllService(data)
		}
	case msg.ServiceDiscoveryCheckService:
		{
			s.CheckService(data)
		}

	}

}

// 注册服务
func (s *Service) RegisterService(data proto.IDataPack) {

	//fmt.Println("收到线程1")
	info := &proto.Discovery{}
	err := data.GetData().ProtoBuf(info)

	if err != nil {
		fmt.Println(data.GetId(), "服务注册解析失败")
		return
	}

	coll := s.IManage.DB().Mongo().GetCollection(utils.ServiceDiscoveryTable)

	dbinfo := &proto.Discovery{}
	err = coll.Find(bson.M{"id": info.Id}).One(&dbinfo)
	if err != nil {
		dbinfo.State = msg.ServiceDiscoveryState
		info.State = msg.ServiceDiscoveryState
		err = coll.Insert(info)
		fmt.Println("添加服务: ", info.Id, info.Topic, info.ServerId)
		s.IManage.Message().Kafka().Send().Async(msg.ServiceDiscoveryListenTopic, s.IManage.Server().GetId(), data.GetRawData())

	}

	if dbinfo.State != msg.ServiceDiscoveryState {
		fmt.Println(dbinfo.Id, " 服务目前处于关闭状态")
		return
	}

	coll2 := s.IManage.DB().Mongo().GetCollection(utils.ServiceDiscoveryTable + info.Topic)
	err = coll2.Find(bson.M{"serverid": info.ServerId}).One(&dbinfo)

	if err == nil {
		fmt.Println(dbinfo.ServerId, "  已存在")
		return
	}
	err = coll2.Insert(info)
	fmt.Println("注册服务: ", info.Id, info.Topic, info.ServerId)
}

// 删除服务
func (s *Service) LogoutService(data proto.IDataPack) {

	info := &proto.Discovery{}
	err := data.GetData().ProtoBuf(info)

	if err != nil {
		fmt.Println(data.GetId(), "服务注册解析失败")
		return
	}
	dbinfo := &proto.Discovery{}
	coll := s.IManage.DB().Mongo().GetCollection(utils.ServiceDiscoveryTable + info.Topic)
	err = coll.Find(bson.M{"serverid": data.GetServerId()}).One(dbinfo)

	fmt.Println(dbinfo)
	if err != nil {
		fmt.Println(data.GetServerId(), "服务数据库删除失败，无数据")
		//return
	}
	err = coll.Remove(dbinfo)
	if err != nil {
		fmt.Println(data.GetServerId(), "服务数据库删除失败")
		//return
	}

	s.IManage.Message().Kafka().Send().Sync(msg.ServiceDiscoveryListenTopic, s.IManage.Server().GetId(), data.GetRawData())

	var allInfo []proto.Discovery

	err = coll.Find(bson.M{}).All(&allInfo)

	//fmt.Println("开始查询", len(allInfo))
	if len(allInfo) == 0 {
		//	fmt.Println("查询服务无数据")
		//s.IManage.Message().Kafka().Send().Async(utils.ServiceDiscoveryListenTopic, s.IManage.Server().GetId(), data.GetRawData())

		coll2 := s.IManage.DB().Mongo().GetCollection(utils.ServiceDiscoveryTable)
		coll2.Find(bson.M{"id": dbinfo.Id}).One(dbinfo)
		coll2.Remove(dbinfo)

		//	fmt.Println("查询服务无数据")

		dbinfo.State = 0
		b := s.IManage.Message().DataPack().Pack(msg.ServiceDiscoveryID, msg.ServiceDiscoveryCloseService, s.IManage.Server().GetId(),
			s.IManage.Server().GetId(), s.IManage.Tool().Protobuf().Encode(dbinfo))

		s.IManage.Message().Kafka().Send().Async(msg.ServiceDiscoveryListenTopic, s.IManage.Server().GetId(), b)

		//fmt.Println("查询服务无数据")

	}

	//	fmt.Println("结束查询")

}

// 查询健康服务
func (s *Service) CheckService(data proto.IDataPack) {
	fmt.Println("收到线程3")

}

// 查询所有健康服务
func (s *Service) CheckAllService(data proto.IDataPack) {

	fmt.Println("收到线程4")

	var dbInfo []proto.Discovery

	coll := s.IManage.DB().Mongo().GetCollection(utils.ServiceDiscoveryTable)

	err := coll.Find(bson.M{}).All(&dbInfo)
	fmt.Println("收到线程4")

	if err != nil {
		fmt.Println("查询服务无数据")
		return
	}
	fmt.Println("收到线程4", data.GetServerId())

	for i := 0; i < len(dbInfo); i++ {
		if dbInfo[i].State == msg.ServiceDiscoveryState {
			pushdata := s.IManage.Tool().Protobuf().Encode(&dbInfo[i])
			b := s.IManage.Message().DataPack().Pack(msg.ServiceDiscoveryID, msg.ServiceDiscoveryRegister, "",
				s.IManage.Server().GetId(), pushdata)
			s.IManage.Message().Kafka().Send().Async(data.GetServerId(), s.IManage.Server().GetId(), b)
			fmt.Println("循环第 ", i+1)
		}

	}
	//	fmt.Println(len(dbInfo))
	//	fmt.Println(dbInfo)

	/*

		//var pushInfo []pd.Discovery
		//d:= make(map[uint32]string)

		push:=make(map[uint32]*pd.Discovery)
		var k uint32
		for i:=0;i<= len(dbInfo);i++{
			if push[dbInfo[i].Id].Id != 0{
			//	d[dbInfo[i].Id]=dbInfo[i].Topic
				push[k]=dbInfo[i]
			}
		}

		fmt.Println(push)


	*/

}
