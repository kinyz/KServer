package main

import (
	"KServer/library/utils"
	proto2 "KServer/proto"
	obj "KServer/server/utils"
	pd2 "KServer/server/utils/pd"
	"fmt"

	"log"
	"net/url"
	"os"
	"os/signal"

	// "strconv"

	"time"

	"github.com/gorilla/websocket"
)

var proto utils.Protobuf
var max = 10

func main() {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	//	Path: "/echo"
	u := url.URL{Scheme: "ws", Host: "127.0.0.1:8889"}
	log.Println("connecting to ", u.String())

	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial err:", err)
		return
	}
	defer conn.Close()

	go timeWriter(conn)

	i := 0
	for {
		//第一个包
		// log.Println("main ReadMessage start")
		_, msg, err := conn.ReadMessage()
		fmt.Println(string(msg))
		if err != nil {
			log.Fatal("read err:", err)
			return
		}
		//log.Println("main ReadMessage read server mt:", mt, " message:", string(msg[:]))
		fmt.Println(string(msg))
		//第三个包主动报错
		if i == 3 {
			panic("主动报错")
		}
		// break
		i++
		// if i > max {
		// 	break
		// }
	}
	//阻塞
	// select {}
	time.Sleep(60 * time.Second)
	log.Println("client exit")
}

func timeWriter(conn *websocket.Conn) {
	var i = 0
	for {
		// log.Println("WriteMessage start timeWriter i = ", i)

		user := &pd2.Account{
			UUID:    "cab02938-4a6a-4e50-b393-94da981e6660",
			Account: "123",
			Token:   "ea456525075570667b1cccaf99356ad0",
			Online:  0,
			State:   0,
		}
		data := proto.Encode(user)
		//发第一个消息
		msg := &proto2.Message{
			Id:       obj.OauthId,
			MsgId:    obj.OauthAccount,
			ClientId: "cab02938-4a6a-4e50-b393-94da981e6660",
			ServerId: "",
			Data:     data,
		}
		pushdata := proto.Encode(msg)
		//	msg := &message.Account{Name: "第一个包 hello,张三", Age: i, Passwd: "123456"}

		conn.WriteMessage(websocket.BinaryMessage, pushdata)

		// //发第二个消息
		// msg = &message.Account{Name: "第二个包 hello, 李四", Age: i, Passwd: "654321"}
		// jsonData, err = json.Marshal(msg)
		// if err != nil {
		// 	log.Println("client timeWriter Marshal err:", err, " msg:", msg)
		// 	break
		// }
		// conn.WriteMessage(websocket.TextMessage, jsonData)

		// // //第三个是回写数据
		// repeatMsg := []byte("第三个包repeat message i = " + strconv.Itoa(i))
		// conn.WriteMessage(websocket.TextMessage, repeatMsg)

		//cpu阻塞下，等待读取完
		time.Sleep(5 * time.Second)

		i++
		if i > max {
			break
		}
	}
	select {}

}
