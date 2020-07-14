package pack

import (
	"KServer/library/iface/ikafka"
	"KServer/library/iface/ilock"
	"KServer/library/iface/iredis"
	"KServer/library/iface/iutils"
	"KServer/library/lock"
	"KServer/proto"
	"KServer/server/utils/msg"

	"fmt"
	"time"
)

type ILockPack interface {
	/*
		加锁
		key string 所需要加锁的Key
		sendKafka bool 应用于已开启kafka处理
		成功返回true
		失败返回false
	*/
	Lock(key string) bool
	/*
		解锁
		key string 所需要加锁的Key
		sendKafka bool 应用于已开启kafka处理
		成功返回true
		失败返回false
	*/
	UnLock(key string) bool
	/*
		自动锁
		key string 所需要加锁的Key
		timeOut int64 设置秒数后自动解锁
		sendKafka bool 解锁失败后应用于已开启kafka处理
		成功返回true
		失败返回false
		自动解锁失败会返回日志
	*/
	AutoLock(key string, timeOut int) bool
	/*
		时间锁
		key string 所需要加锁的Key
		timeSleep int64 毫秒 设置毫秒进行一次加锁
		timeOut int64 毫秒 设置毫秒加锁返回结果
		sendKafka bool 应用于已开启kafka处理
		成功返回true
		失败返回false
	*/
	TimeLock(key string, timeSleep int, timeOut int) bool
	/*
		次数锁
		key string 所需要加锁的Key
		num int 设置次数内锁
		sendKafka bool 应用于已开启kafka处理
		成功返回true
		失败返回false
	*/
	NumLock(key string, num int) bool

	// 检查LockKey是否已存在
	Check(key string) bool

	// 将LockKey发送到kafka进行处理 type自定义执行的动作
	SendKafka(key string, Type uint32)
}

type LockPack struct {
	send      ikafka.ISend
	lock      ilock.ILock
	lockTopic string
	serverId  string
	dataPack  proto.IDataPack
	protoTool iutils.IProtobuf
}

func NewILockPack(serverId string, lockTopic string, dataPack proto.IDataPack, protoTool iutils.IProtobuf, redis iredis.IRedisPool, send ikafka.ISend) ILockPack {
	return &LockPack{serverId: serverId, lockTopic: lockTopic, lock: lock.NewILock(redis), send: send, dataPack: dataPack, protoTool: protoTool}
}

// 锁Key
func (l *LockPack) Lock(key string) bool {
	if l.lock.Lock(key) {
		return true
	}
	return false
}

// 解锁Key
func (l *LockPack) UnLock(key string) bool {
	if l.lock.UnLock(key) {
		return true
	}
	return false
}

// 自动锁 设置多少秒 自动解锁
func (l *LockPack) AutoLock(key string, timeOut int) bool {
	if l.lock.Lock(key) {
		go func() {
			//	fmt.Println("自动锁开启")
			time.Sleep(time.Duration(timeOut) * time.Second)
			if !l.UnLock(key) {
				fmt.Println(key, "自动解锁失败")
				return
			}
		}()
		return true
	}
	return false
}

// 时间锁Key
func (l *LockPack) TimeLock(key string, timeSleep int, timeOut int) bool {
	timeout := time.After(time.Millisecond * time.Duration(timeOut))
	finish := make(chan bool)
	count := 0
	go func() {
		for {
			select {
			case <-timeout:
				//	fmt.Println("超时")
				finish <- true
				return
			default:
				fmt.Println("正在执行加锁")
				if l.Lock(key) {
					//	fmt.Println("加锁成功")
					count++
					finish <- true
					return
				}

			}
			time.Sleep(time.Millisecond * time.Duration(timeSleep))
		}
	}()
	<-finish
	if count != 0 {
		return true
	}
	return false

}

// 次数锁 设置次数锁
func (l *LockPack) NumLock(key string, num int) bool {

	for i := 0; i < num; i++ {
		if l.Lock(key) {
			return true
		}
		fmt.Println("次数加锁", i)
	}
	return false
}

func (l *LockPack) SendKafka(key string, Type uint32) {
	lockMsg := &proto.LockMessage{
		Key:  key,
		Type: Type,
	}
	l.send.Async(l.lockTopic, l.serverId, l.dataPack.Pack(msg.LockId, msg.LockTypeMsgId, l.serverId, l.serverId, l.protoTool.Encode(lockMsg)))

}

func (l *LockPack) Check(key string) bool {
	return l.lock.Check(key)
}
