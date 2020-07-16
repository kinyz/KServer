package lock

import (
	"KServer/library/kiface/ilock"
	"KServer/library/kiface/iredis"
	"fmt"
)

type UnLock struct {
	IRedis   iredis.IRedisPool
	LockText string
}

func NewIUnLock(LockText string, redis iredis.IRedisPool) ilock.IUnLock {
	return &UnLock{IRedis: redis, LockText: LockText}
}

func (ul *UnLock) Do(Key string) bool {
	//conn:=l.IRedis.GetMasterConn()
	_, err := ul.IRedis.GetMasterConn().Del(Key)
	if err != nil {
		fmt.Println("UnLock ", Key, "Error :", err)
		return false
	}
	return true
}
