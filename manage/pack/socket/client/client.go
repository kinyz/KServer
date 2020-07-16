package client

import (
	"KServer/library/kiface/isocket"
	socket2 "KServer/manage/socket"

	"KServer/manage/config"
	"fmt"
	"time"
)

type IClientPack interface {
	/*
		添加Client
		uuid string 所需要加锁的Key
		token string 应用于已开启kafka处理
		conn isocket.IConnection
	*/
	AddClient(uuid string, token string, conn isocket.IConnection)
	/*
		移除Client
		id uint32 一般用connId
	*/
	Remove(id uint32)
	/*
		获取Client
		uuid string
		返回IClient对象
	*/
	GetClient(uuid string) socket2.IClient
	/*
		获取Client状态
		uuid string
		返回值 bool
	*/
	GetState(uuid string) bool
	/*
		获取当前Client数量
		返回值 int
	*/
	GetOnlineNum() int
	/*
		通过clientId 查询UUID
		id uint32
		返回值 string
	*/
	GetIdByConnId(id uint32) string
	/*
		通知所有Client
		data []byte
	*/
	SendAll(data []byte)
	/*
		设置Client关闭回调
		uuid string
		fun func()
	*/
	SetClose(uuid string, fun func())
}

type ClientPack struct {
	ConnId map[uint32]string
	Client map[string]socket2.IClient
	config *config.ManageConfig
	close  map[string]func()
}

func NewIClientPack() IClientPack {
	return &ClientPack{
		Client: make(map[string]socket2.IClient),
		ConnId: make(map[uint32]string),
		close:  make(map[string]func()),
	}

}
func (cp *ClientPack) AddClient(uuid string, token string, conn isocket.IConnection) {
	c := socket2.NewClient(conn, token)
	cp.Client[uuid] = c
	cp.ConnId[conn.GetConnID()] = uuid
}

func (cp *ClientPack) Remove(id uint32) {

	delete(cp.Client, cp.GetIdByConnId(id))
	delete(cp.ConnId, id)

}

func (cp *ClientPack) GetClient(uuid string) socket2.IClient {
	if cp.GetState(uuid) {
		return cp.Client[uuid]
	}
	return nil
}

func (cp *ClientPack) GetState(uuid string) bool {
	if cp.Client[uuid] != nil {
		return true
	}
	return false
}

func (cp *ClientPack) GetOnlineNum() int {
	return len(cp.Client)
}

// 通知全部客户端
func (cp *ClientPack) SendAll(data []byte) {

	for _, uuid := range cp.Client {
		if cp.GetState(cp.GetIdByConnId(uuid.GetConnId())) {
			err := uuid.SendBuff(data)
			if err != nil {
				fmt.Println("全局消息发送失败", uuid)
			}
		}
	}
	//return len(cp.Client)
}

func (cp *ClientPack) GetIdByConnId(id uint32) string {
	return cp.ConnId[id]
}

func (cp *ClientPack) SetClose(uuid string, fun func()) {
	if cp.close[uuid] == nil {
		cp.close[uuid] = fun
		cp.SetCloseTimer(uuid)
		//fmt.Println("成功设置关闭回调")
	}

}

func (cp *ClientPack) SetCloseTimer(uuid string) {
	go func() {
		for {
			//fmt.Println("正在执行定时器")
			if !cp.GetState(uuid) {
				//	fmt.Println("发现掉线的客户端")
				if cp.close[uuid] != nil {
					//fmt.Println("正在移除监听")
					go func() {
						cp.close[uuid]()
						delete(cp.close, uuid)
					}()
					break
				}

				//cp.close[k]()
			}
			//	fmt.Println(cp.close)

			time.Sleep(5 * time.Second)
		}

	}()
}
