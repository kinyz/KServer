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
	u := url.URL{Scheme: "ws", Host: "127.0.0.1:9999"}
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

	time.Sleep(10 * time.Second)
	log.Println("client exit")
}

func timeWriter(conn *websocket.Conn) {

	// log.Println("WriteMessage start timeWriter i = ", i)

	//imsg := utils.NewIDataPack()
	user := &pd2.Account{
		UUID:    "cab02938-4a6a-4e50-b393-94da981e6660",
		Account: "123",
		Token:   "c84e59cadf42df0361ea28ba45366758",
		Online:  0,
		State:   0,
	}

	v := proto.Encode(user)

	data := proto2.NewIMessage(obj.OauthId, obj.OauthAccount, "cab02938-4a6a-4e50-b393-94da981e6660", "", v)

	conn.WriteMessage(websocket.BinaryMessage, data)

	time.Sleep(5 * time.Second)

	//select {}

}
