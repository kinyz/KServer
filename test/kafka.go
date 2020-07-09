package main

import (
	"KServer/library/iface"
	"KServer/library/kafka"
)

func main() {

	t := &Test{}
	msg := kafka.NewIMessage()
	msg.Router().AddRouter("test_go", t)

	msg.Router().StartListen([]string{"140.143.247.121:30557"}, "ijioj", -1)

	select {}
	//kafka.Close()

	//i:=kafka.NewIConsumer()
	//i.NewConsumerGroup([]string{"140.143.247.121:30557"},"cew",-1)

}

type Test struct {
}

func (t *Test) ResponseHandle(response iface.IResponse) {

}
