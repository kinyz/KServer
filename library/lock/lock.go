package lock

import (
	"KServer/library/iface/ilock"
	"KServer/library/iface/iredis"
	"fmt"
)

type Lock struct {
	IRedis iredis.IRedisPool
}

func NewILock(redis iredis.IRedisPool) ilock.ILock {
	return &Lock{IRedis: redis}
}

func (l *Lock) Lock(Key string) bool {

	v, err := l.IRedis.GetSlaveConn().Get(Key).Value()

	if err != nil || v != nil {
		return false
	}

	_, err = l.IRedis.GetMasterConn().Set(Key).String("true")
	if err != nil {
		fmt.Println("Lock ", Key, "Error :", err)
		return false
	}
	return true
}

func (l *Lock) UnLock(Key string) bool {
	//conn:=l.IRedis.GetMasterConn()
	_, err := l.IRedis.GetMasterConn().Del(Key)
	if err != nil {
		fmt.Println("UnLock ", Key, "Error :", err)
		return false
	}
	return true
}

func (l *Lock) Check(Key string) bool {

	v, err := l.IRedis.GetSlaveConn().Get(Key).Value()

	if err != nil {
		return false
	}
	if v != nil {
		return true
	}
	return false
}
