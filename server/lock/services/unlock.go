package services

import (
	"KServer/manage"
	"KServer/proto"
	"KServer/server/utils/msg"
	"fmt"
	"time"
)

type UnLock struct {
	m manage.IManage
}

func NewUnLock(m manage.IManage) *UnLock {
	return &UnLock{m: m}
}

func (u *UnLock) UnlockHandle(data proto.IDataPack) {

	//fmt.Println("收到请求",data.GetMsgId())
	switch data.GetMsgId() {
	case msg.LockTypeMsgId:
		u.LockType(data)

	}
}

func (u *UnLock) LockType(data proto.IDataPack) {
	lockInfo := &proto.LockMessage{}
	data.GetData().ProtoBuf(lockInfo)

	switch lockInfo.Type {
	case msg.LockTypeAutoUnLock:
		u.AutoUnLock(lockInfo)
	}

}
func (u *UnLock) AutoUnLock(lockInfo *proto.LockMessage) {

	go func() {
		for i := 0; i < 20; i++ {
			fmt.Println("收到自动解锁请求", lockInfo)
			time.Sleep(10 * time.Second)
			if u.m.Lock().UnLock().Do(lockInfo.Key) {
				fmt.Println("解锁完成", lockInfo)

				break
			}
		}
	}()

}
