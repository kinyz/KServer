package pack

import (
	iface "KServer/library/iface/ikafka"
	"KServer/library/kafka"
	"KServer/server/utils"
)

type IKafkaPack interface {
	DataPack() utils.IDataPack
	AddRouter(topic string, id uint32, response func(data utils.IDataPack))
	Send() iface.ISend
	ResponseHandle(req iface.IResponse)
	StartListen(addr []string, group string, offset int64) func()
	StartOtherListen(topic []string, addr []string, group string, offset int64) func()
}

type KafkaPack struct {
	IKafka    iface.IKafka
	topic     map[string]map[uint32]func(data utils.IDataPack)
	IDataPack utils.IDataPack
}

func NewKafkaPack() IKafkaPack {

	//map2 := map1["error"][0]
	return &KafkaPack{
		IKafka:    kafka.NewIKafka(),
		topic:     make(map[string]map[uint32]func(data utils.IDataPack)),
		IDataPack: utils.NewIDataPack(),
	}
}
func (m *KafkaPack) AddRouter(topic string, id uint32, response func(data utils.IDataPack)) {

	if m.topic[topic] == nil {
		m.topic[topic] = make(map[uint32]func(data utils.IDataPack))
	}
	m.topic[topic][id] = response
	m.IKafka.Router().AddRouter(topic, m)
	//fmt.Println(m.topic)
}

func (m *KafkaPack) Send() iface.ISend {
	return m.IKafka.Send()
}
func (m *KafkaPack) ResponseHandle(req iface.IResponse) {
	_ = m.IDataPack.UnPack(req.GetData().Bytes())
	//fmt.Println(req.GetTopic())
	if m.topic[req.GetTopic()][m.IDataPack.GetId()] != nil {
		m.topic[req.GetTopic()][m.IDataPack.GetId()](m.IDataPack)
	}
}

func (m *KafkaPack) DataPack() utils.IDataPack {
	return m.IDataPack
}

func (m *KafkaPack) StartListen(addr []string, group string, offset int64) func() {
	return m.IKafka.Router().StartListen(addr, group, offset)
}
func (m *KafkaPack) StartOtherListen(topic []string, addr []string, group string, offset int64) func() {
	return m.IKafka.Router().StartOtherListen(topic, addr, group, offset)
}
