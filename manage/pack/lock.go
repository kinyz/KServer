package pack

import (
	"KServer/library/kiface/ikafka"
	"KServer/library/kiface/ilock"
	"KServer/library/kiface/iredis"
	"KServer/library/kiface/iutils"
	"KServer/library/lock"
	"KServer/proto"
	"KServer/server/utils/msg"
)

type ILockPack interface {
	Lock() ilock.ILock

	UnLock() ilock.IUnLock

	// 将LockKey发送到kafka进行处理 type自定义执行的动作
	SendKafka(key string, Type uint32)
}

type LockPack struct {
	send      ikafka.ISend
	lock      ilock.ILock
	unlock    ilock.IUnLock
	lockTopic string
	serverId  string
	dataPack  proto.IDataPack
	protoTool iutils.IProtobuf
}

func NewILockPack(serverId string, lockTopic string, dataPack proto.IDataPack, protoTool iutils.IProtobuf, redis iredis.IRedisPool, send ikafka.ISend) ILockPack {
	return &LockPack{serverId: serverId, lockTopic: lockTopic, unlock: lock.NewIUnLock(serverId, redis), lock: lock.NewILock(serverId, redis), send: send, dataPack: dataPack, protoTool: protoTool}
}

// 锁Key
func (l *LockPack) Lock() ilock.ILock {

	return l.lock
}

// 解锁Key
func (l *LockPack) UnLock() ilock.IUnLock {

	return l.unlock
}

func (l *LockPack) SendKafka(key string, Type uint32) {
	lockMsg := &proto.LockMessage{
		Key:  key,
		Type: Type,
	}
	l.send.Async(l.lockTopic, l.serverId, l.dataPack.Pack(msg.LockId, msg.LockTypeMsgId, l.serverId, l.serverId, l.protoTool.Encode(lockMsg)))
}
