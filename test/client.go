package main

import (
	"KServer/library/socket/znet"
	tool2 "KServer/library/utils"
	"KServer/server/utils"
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

	conn, err := net.Dial("tcp", "127.0.0.1:8889")
	if err != nil {
		fmt.Println("client start err, exit!")
		return
	}
	imsg := utils.NewIDataPack()
	user := &pb.Account{
		UUID:    "27c340b1-6d1b-4893-a14c-abb1f81829c4",
		Account: "116175894",
		Token:   "123ebf90eb9f79be7ed1baaac6704617",
		Online:  0,
		State:   0,
	}

	dp := znet.NewDataPack()

	v := pd.Encode(user)

	msg, _ := dp.Pack(znet.NewMsgPackage(utils.OauthMsgId, imsg.Pack(utils.OauthMsgId, user.UUID, "1", 0, utils.OauthAccount, v)))
	//fmt.Println(msg)
	//for i := 0; i < 5; i++ {
	_, err = conn.Write(msg)

	//}
	if err != nil {
		fmt.Println("client write err: ", err)
		return
	}

	for {

		//先读出流中的head部分
		headData := make([]byte, dp.GetHeadLen())
		_, err := io.ReadFull(conn, headData)
		if err != nil {
			fmt.Println("client read head err: ", err)
			return
		}

		// 将headData字节流 拆包到msg中
		msgHead, err := dp.Unpack(headData)
		if err != nil {
			fmt.Println("client unpack head err: ", err)
			return
		}

		if msgHead.GetDataLen() > 0 {
			//msg 是有data数据的，需要再次读取data数据
			msg := msgHead.(*znet.Message)
			msg.Data = make([]byte, msg.GetDataLen())

			//根据dataLen从io中读取字节流
			_, err := io.ReadFull(conn, msg.Data)
			if err != nil {
				fmt.Println("client unpack data err")
				return
			}
			if imsg.UnPack(msg.Data) == nil {
				fmt.Println(imsg.GetClientId())

				fmt.Println(imsg.GetMsgId())

				fmt.Println(imsg.GetClientConnId())

				fmt.Println(imsg.GetServerId())

				fmt.Println(imsg.GetDate().String())
				acc := &pb.Account{}
				err := imsg.GetDate().ProtoBuf(acc)
				if err != nil {
					fmt.Println("err=", err)
				}
				fmt.Println("acc=" + acc.Account)
			}

			fmt.Printf("==> Client receive Msg: Id = %d, len = %d , data = %s\n", msg.Id, msg.DataLen, msg.Data)
		}

		time.Sleep(time.Second)
	}
}
