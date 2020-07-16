package lock

import (
	"KServer/library/kiface/ilock"
	"KServer/library/kiface/iredis"
	"KServer/library/kiface/iutils"
	"KServer/library/utils"
	"fmt"
	"time"
)

type Lock struct {
	IRedis   iredis.IRedisPool
	LockText string
	QueueId  int64
	pdtool   iutils.IProtobuf
}

func NewILock(LockText string, redis iredis.IRedisPool) ilock.ILock {

	return &Lock{IRedis: redis, pdtool: utils.NewIProtobuf(), LockText: LockText}
}

func (l *Lock) Do(Key string) bool {
	v, err := l.IRedis.GetSlaveConn().Get(Key).Value()
	if err != nil || v != nil {
		return false
	}
	_, err = l.IRedis.GetMasterConn().Set(Key).String(l.LockText)
	if err != nil {
		fmt.Println("Lock ", Key, "Error :", err)
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

// 自动锁 设置多少秒 自动解锁
func (l *Lock) Auto(key string, timeOut int) bool {
	if l.Do(key) {
		go func() {
			//	fmt.Println("自动锁开启")
			time.Sleep(time.Duration(timeOut) * time.Second)
			_, err := l.IRedis.GetMasterConn().Del(key)
			if err != nil {
				fmt.Println("AutoUnLock ", key, " Error :", err)
			}

		}()
		return true
	}
	return false
}

// 时间锁Key
func (l *Lock) Time(key string, timeSleep int, timeOut int) bool {
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
				if l.Do(key) {
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
func (l *Lock) Num(key string, num int) bool {

	for i := 0; i < num; i++ {
		if l.Do(key) {
			return true
		}
		//fmt.Println("次数加锁", i)
	}
	return false
}

func (l *Lock) Queue(key string, timeOut int64) ilock.IQueue {
	l.QueueId++
	return NewILockQueue(key, l.QueueId, l.LockText, timeOut, l.pdtool, l.IRedis)
}
