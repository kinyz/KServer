package pack

import (
	"KServer/library/iface/ikafka"
	"KServer/library/iface/ilock"
	"KServer/library/iface/iredis"
	"KServer/library/lock"
)

type ILockPack interface {
	Lock(key string) bool
	UnLock(key string) bool
}

type LockPack struct {
	send      ikafka.ISend
	lock      ilock.ILock
	lockTopic string
}

func NewILockPack(lockTopic string, redis iredis.IRedisPool, send ikafka.ISend) ILockPack {
	return &LockPack{lockTopic: lockTopic, lock: lock.NewILock(redis), send: send}
}

// 锁Key
func (l *LockPack) Lock(key string) bool {
	if l.lock.Lock(key) {
		l.send.Async(l.lockTopic, key, []byte("lock"))
		return true
	}
	return false
}

// 解锁Key
func (l *LockPack) UnLock(key string) bool {
	if l.lock.UnLock(key) {
		return true
	}
	l.send.Async(l.lockTopic, key, []byte("unlock"))
	return false
}
