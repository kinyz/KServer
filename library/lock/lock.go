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
	fmt.Println("获取锁1")

	v, err := l.IRedis.GetSlaveConn().Get(Key).Value()
	fmt.Println("获取锁2", v)

	if err != nil || v != nil {
		return false
	}

	fmt.Println("获取锁3")
	_, err = l.IRedis.GetMasterConn().Set(Key).String("true")
	if err != nil {
		fmt.Println("获取锁14")

		fmt.Println("Lock ", Key, "Error :", err)
		return false
	}
	fmt.Println("获取锁5")

	fmt.Println("获取锁5")

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
