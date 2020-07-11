package client

import (
	"KServer/library/iface/isocket"
	"KServer/manage/client"
	"KServer/manage/config"
	"fmt"
	"time"
)

type IClientPack interface {
	AddClient(uuid string, token string, conn isocket.IConnection)
	Remove(id uint32)
	GetClient(uuid string) client.IClient
	GetState(uuid string) bool
	GetOnlineNum() int
	GetIdByConnId(id uint32) string
	SetClose(uuid string, fun func())
}

type ClientPack struct {
	ConnId map[uint32]string
	Client map[string]client.IClient
	config *config.ManageConfig
	close  map[string]func()
}

func NewIClientPack() IClientPack {
	return &ClientPack{
		Client: make(map[string]client.IClient),
		ConnId: make(map[uint32]string),
		close:  make(map[string]func()),
	}

}
func (cp *ClientPack) AddClient(uuid string, token string, conn isocket.IConnection) {
	c := client.NewClient(conn, token)
	cp.Client[uuid] = c
	cp.ConnId[conn.GetConnID()] = uuid
}

func (cp *ClientPack) Remove(id uint32) {

	delete(cp.Client, cp.GetIdByConnId(id))
	delete(cp.ConnId, id)

}

func (cp *ClientPack) GetClient(uuid string) client.IClient {
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

func (cp *ClientPack) GetIdByConnId(id uint32) string {
	return cp.ConnId[id]
}

func (cp *ClientPack) SetClose(uuid string, fun func()) {
	if cp.close[uuid] == nil {
		cp.close[uuid] = fun
		cp.SetCloseTimer(uuid)
		fmt.Println("成功设置关闭回调")
	}

}

func (cp *ClientPack) SetCloseTimer(uuid string) {
	go func() {
		for {
			fmt.Println("正在执行定时器")
			if !cp.GetState(uuid) {
				fmt.Println("发现掉线的客户端")
				if cp.close[uuid] != nil {
					fmt.Println("正在移除监听")
					cp.close[uuid]()
					delete(cp.close, uuid)
					break
				}

				//cp.close[k]()
			}
			fmt.Println(cp.close)

			time.Sleep(5 * time.Second)
		}

	}()
}
