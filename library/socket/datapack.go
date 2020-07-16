package socket

import (
	"KServer/library/kiface/isocket"
	"KServer/library/socket/utils"
	"bytes"
	"encoding/binary"
	"errors"
)

//封包拆包类实例，暂时不需要成员
type DataPack struct{}

//封包拆包实例初始化方法
func NewDataPack() *DataPack {
	return &DataPack{}
}

//获取包头长度方法
func (dp *DataPack) GetHeadLen() uint32 {
	// DataLen uint32(4字节)
	return 4
}

//封包方法(压缩数据)
func (dp *DataPack) Pack(data []byte) ([]byte, error) {
	//创建一个存放bytes字节的缓冲

	dataBuff := bytes.NewBuffer([]byte{})

	//fmt.Println(uint32(len(data)))
	//fmt.Println(dataBuff)
	//data:=proto.Encode(msg)

	len := uint32(len(data))
	//写dataLen
	if err := binary.Write(dataBuff, binary.LittleEndian, len); err != nil {
		return nil, err
	}
	//fmt.Println("data长度 ", msg.GetDataLen())
	//fmt.Println(dataBuff.Bytes())
	//写data数据
	if err := binary.Write(dataBuff, binary.LittleEndian, data); err != nil {
		return nil, err
	}
	//	fmt.Println(dataBuff.Bytes())
	return dataBuff.Bytes(), nil

}

//拆包方法(解压数据) 获取数据的长度
func (dp *DataPack) Unpack(binaryData []byte) (isocket.IMessage, error) {
	//创建一个从输入二进制数据的ioReader

	dataBuff := bytes.NewReader(binaryData)

	msg := &Message{}
	//读dataLen
	if err := binary.Read(dataBuff, binary.LittleEndian, &msg.DataLen); err != nil {
		return nil, err
	}

	//fmt.Println(msg.DataLen)
	//判断dataLen的长度是否超出我们允许的最大包长度
	if utils.GlobalObject.MaxPacketSize > 0 && msg.DataLen > utils.GlobalObject.MaxPacketSize {
		return nil, errors.New("too large msg data received")
	}

	//d:=proto.NewMsgNull()
	//tool.Decode(binaryData,d)
	//msg := &Message{d}
	//fmt.Println(msg.GetId(),msg.GetData())

	//这里只需要把head的数据拆包出来就可以了，然后再通过head的长度，再从conn读取一次数据
	return msg, nil
}
