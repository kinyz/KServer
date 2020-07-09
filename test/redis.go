package main

import (
	"KServer/library/proto"
	"KServer/library/redis"
	tool2 "KServer/library/utils"
	"fmt"
)

var tool tool2.Protobuf

func main() {

	conf := redis.GetRedisConf("/conf/redis.yaml")

	r := redis.NewRedisPool()
	r.StartMasterPool(conf.Master.Host+":"+conf.Master.Port, conf.Master.PassWord, conf.Master.MaxIdle, conf.Master.MaxActive)
	r.StartSlavePool(conf.Slave.Host+":"+conf.Slave.Port, conf.Slave.PassWord, conf.Slave.MaxIdle, conf.Slave.MaxActive)
	//fmt.Println(redis.SlaveHost+":"+redis.Port,redis.SlavePassWord)
	msg := proto.NewMsg(100, []byte("我是结构体"))
	r.GetMasterConn().Set("key1").String("string")
	r.GetMasterConn().Set("key2").Bytes([]byte("byte"))
	r.GetMasterConn().Set("key3").Json(msg)
	r.GetMasterConn().Set("key4").ProtoBuf(msg)

	fmt.Println("key1=", r.GetSlaveConn().Get("key1").String())
	fmt.Println("key2=", r.GetSlaveConn().Get("key2").Bytes())
	json := proto.NewMsgNull()
	r.GetSlaveConn().Get("key3").Json(json)
	fmt.Println("key3=", json.Id, json.DataLen, string(json.Data))
	protovalue := proto.NewMsgNull()
	r.GetSlaveConn().Get("key4").ProtoBuf(protovalue)
	fmt.Println("key4=", protovalue.Id, protovalue.DataLen, protovalue.Data)

	for {
		fmt.Println("key2=", r.GetSlaveConn().Get("key2").Bytes())
	}
	/*
		redis.NewPool()
		redis.NewSlavePool()
		_, _ = redis.Check("system") // 进行一次redis健康检查

		msg := &proto.Msg{}
		msg.Id = 100
		msg.Data = []byte("DoConnection BEGIN...")

		//_, err := redis.SetValueByProto("proto", msg)

		k, err := redis.PushValueByProto("listkey", msg)
		fmt.Println(k)
		if err != nil {
			fmt.Println(err)
		}

		msg2 := &proto.Msg{}

		err = redis.GetValueToProto("proto", msg2)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(msg2.Id, string(msg2.Data))
	*/
}
