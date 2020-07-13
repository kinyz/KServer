package kafkaPack

import (
	"KServer/library/iface/ikafka"
	"KServer/library/iface/iutils"
	"KServer/library/kafka"
	utils2 "KServer/library/utils"
	"KServer/manage/discover/pd"
	"KServer/proto"
	"KServer/server/utils"
)

type IKafkaPack interface {
	DataPack() proto.IDataPack
	AddRouter(topic string, id uint32, response func(data proto.IDataPack))
	Send() ikafka.ISend
	ResponseHandle(req ikafka.IResponse)
	StartListen(addr []string, group string, offset int64) func()
	StartOtherListen(topic []string, addr []string, group string, offset int64) func()

	// 向服务中心注册一个服务
	CallRegisterService(id uint32, topic string, serverId string, host string, port string, Type string)
	// 向服务中心注销一个服务
	CallLogoutService(id uint32, Topic string, serverId string)
	// 向服务中心关闭一个主线程服务
	CallCloseServiceState(id uint32)
	// 向服务中心打开一个主线程服务
	CallOpenServiceState(id uint32)
	// 查询全部服务
	CallCheckAllService(serverId string)
}

type KafkaPack struct {
	IKafka    ikafka.IKafka
	topic     map[string]map[uint32]func(data proto.IDataPack)
	IDataPack proto.IDataPack
	p         iutils.IProtobuf
}

func NewKafkaPack() IKafkaPack {

	//map2 := map1["error"][0]
	return &KafkaPack{
		IKafka:    kafka.NewIKafka(),
		topic:     make(map[string]map[uint32]func(data proto.IDataPack)),
		IDataPack: proto.NewIDataPack(),
		p:         utils2.NewIProtobuf(),
	}
}
func (m *KafkaPack) AddRouter(topic string, id uint32, response func(data proto.IDataPack)) {

	if m.topic[topic] == nil {
		m.topic[topic] = make(map[uint32]func(data proto.IDataPack))
	}
	m.topic[topic][id] = response
	m.IKafka.Router().AddRouter(topic, m)
	//fmt.Println(m.topic)
}

func (m *KafkaPack) Send() ikafka.ISend {
	return m.IKafka.Send()
}
func (m *KafkaPack) ResponseHandle(req ikafka.IResponse) {
	_ = m.IDataPack.UnPack(req.GetData().Bytes())
	//fmt.Println(req.GetTopic())
	if m.topic[req.GetTopic()][m.IDataPack.GetId()] != nil {
		m.topic[req.GetTopic()][m.IDataPack.GetId()](m.IDataPack)
	}
}

func (m *KafkaPack) DataPack() proto.IDataPack {
	return m.IDataPack
}

func (m *KafkaPack) StartListen(addr []string, group string, offset int64) func() {
	return m.IKafka.Router().StartListen(addr, group, offset)
}
func (m *KafkaPack) StartOtherListen(topic []string, addr []string, group string, offset int64) func() {
	return m.IKafka.Router().StartOtherListen(topic, addr, group, offset)
}

// 向服务中心注册一个服务
func (m *KafkaPack) CallRegisterService(id uint32, topic string, serverId string, host string, port string, Type string) {
	data := &pd.Discovery{
		Id:       id,
		Topic:    topic,
		ServerId: serverId,
		Host:     host,
		Port:     port,
		Type:     Type,
	}
	b := m.DataPack().Pack(utils.ServiceDiscoveryID, utils.ServiceDiscoveryRegister, serverId,
		serverId, m.p.Encode(data))
	m.IKafka.Send().Async(utils.ServiceDiscoveryTopic, serverId, b)
}

// 向服务中心注销一个服务
func (m *KafkaPack) CallLogoutService(id uint32, Topic string, serverId string) {
	data := &pd.Discovery{
		Id:       id,
		Topic:    Topic,
		ServerId: serverId,
		Host:     "",
		Port:     "",
		Type:     "",
	}
	b := m.DataPack().Pack(utils.ServiceDiscoveryID, utils.ServiceDiscoveryLogoutService, serverId,
		serverId, m.p.Encode(data))
	m.IKafka.Send().Async(utils.ServiceDiscoveryTopic, serverId, b)
}

// 向服务中心关闭一个主线程服务
func (m *KafkaPack) CallCloseServiceState(id uint32) {
	data := &pd.Discovery{
		Id:       id,
		Topic:    utils.ServiceDiscoveryTopic,
		ServerId: "",
		Host:     "",
		Port:     "",
		Type:     "kafka",
	}
	b := m.DataPack().Pack(utils.ServiceDiscoveryID, utils.ServiceDiscoveryCloseService, "",
		"", m.p.Encode(data))
	m.IKafka.Send().Async(utils.ServiceDiscoveryTopic, "", b)
}

// 向服务中心关闭一个主线程服务
func (m *KafkaPack) CallOpenServiceState(id uint32) {
	data := &pd.Discovery{
		Id:       id,
		Topic:    utils.ServiceDiscoveryTopic,
		ServerId: "",
		Host:     "",
		Port:     "",
		Type:     "kafka",
	}
	b := m.DataPack().Pack(utils.ServiceDiscoveryID, utils.ServiceDiscoveryLogoutService, "",
		"", m.p.Encode(data))
	m.IKafka.Send().Async(utils.ServiceDiscoveryTopic, "", b)
}

// 向服务中心关闭一个主线程服务
func (m *KafkaPack) CallCheckAllService(serverId string) {

	b := m.DataPack().Pack(utils.ServiceDiscoveryID, utils.ServiceDiscoveryCheckAllService, "",
		serverId, []byte("check all service"))
	m.IKafka.Send().Async(utils.ServiceDiscoveryTopic, "", b)
}
