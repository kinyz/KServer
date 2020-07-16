package main

import (
	"KServer/library/socket"
	tool2 "KServer/library/utils"
	proto2 "KServer/proto"
	msg2 "KServer/server/utils/msg"
	pb "KServer/server/utils/pd"
	"fmt"
	"io"
	"net"

	"time"
)

var pd tool2.Protobuf

func main() {

	ClientTest(3)

}

func ClientTest(i uint32) {

	fmt.Println("Client Test ... start")
	//3秒之后发起测试请求，给服务端开启服务的机会
	//time.Sleep(3 * time.Second)

	conn, err := net.Dial("tcp", "140.143.247.121:32106")
	if err != nil {
		fmt.Println("client start err, exit!")
		return
	}
	//imsg := utils.NewIDataPack()
	user := &pb.Account{
		UUID:    "27c340b1-6d1b-4893-a14c-abb1f81829c4",
		Account: "116175894",
		Token:   "371e8b8f86c88fce4717f034bf9692a8",
		Online:  0,
		State:   0,
	}

	dp := socket.NewDataPack()

	v := pd.Encode(user)

	data := proto2.NewIMessage(msg2.OauthId, msg2.OauthAccount, "b3307e8015acec7d5994ab8564b0a000", "", v)
	fmt.Println(string(data))
	msg, _ := dp.Pack(data)
	//fmt.Println(string(v))
	//for i := 0; i < 5; i++ {
	_, err = conn.Write(msg)

	//}
	if err != nil {
		fmt.Println("client write err: ", err)
		return
	}

	for {

		//读取客户端的Msg head
		headData := make([]byte, dp.GetHeadLen())
		//m:=proto.NewMsgNull()
		//c.Protobuf.Decode()
		if _, err := io.ReadFull(conn, headData); err != nil {
			fmt.Println("read msg head error ", err)
			break
		}
		fmt.Println("head= ", headData)
		msg, err := dp.Unpack(headData)
		if err != nil {
			fmt.Println("unpack error ", err)
			break
		}

		var data []byte
		if msg.GetDataLen() > 0 {
			data = make([]byte, msg.GetDataLen())
			if _, err := io.ReadFull(conn, data); err != nil {
				fmt.Println("read msg data error ", err)
				break
			}
		}
		mess := &proto2.Message{}
		pd.Decode(data, mess)
		msg.SetMessage(mess)

		fmt.Println(mess)
		fmt.Printf("==> Client receive Msg: Id = %d, msgid = %d , data = %s\n", msg.GetId(), msg.GetMsgId(), msg.GetData())
	}

	time.Sleep(time.Second)
}
