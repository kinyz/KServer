package lock

import (
	"KServer/library/kiface/ilock"
	"KServer/library/kiface/iredis"
	"KServer/library/kiface/iutils"
	"KServer/proto"
	"errors"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"time"
)

type Queue struct {
	IRedis   iredis.IRedisPool
	Key      string
	LockText string
	QueueId  int64
	pdtool   iutils.IProtobuf
	timeout  int64
	pushdata []byte
}

func NewILockQueue(Key string, QueueId int64, LockText string, timeout int64, pdtool iutils.IProtobuf, IRedis iredis.IRedisPool) ilock.IQueue {
	return &Queue{Key: Key, QueueId: QueueId, LockText: LockText, timeout: timeout, pdtool: pdtool, IRedis: IRedis}
}

func (q *Queue) Lock(timeSleep int, timeOut int64) error {
	v := &proto.LockQueueMessage{
		Id:      q.QueueId,
		Text:    q.LockText,
		TimeOut: q.timeout,
	}
	conn := q.IRedis.GetRawMasterConn()
	defer conn.Close()
	q.pushdata = q.pdtool.Encode(v)
	i, err := redis.Int(conn.Do("rpush", q.Key, q.pushdata))
	if err != nil {
		return err
	}
	if i == 1 {
		return nil
	}
	timeout := time.After(time.Millisecond * time.Duration(timeOut))
	finish := make(chan bool)
	errtext := ""
	go func() {
		for {
			select {
			case <-timeout:
				err = q.UnLock()
				if err != nil {
					errtext = err.Error()
					finish <- true
					return
				}
				errtext = "timeout"
				finish <- true
				return
			default:
				fmt.Println("正在执行加锁")
				b, err := redis.Bytes(conn.Do("lindex", q.Key, 0))
				if err != nil {
					errtext = err.Error()
					finish <- true
					return
				}
				QMsg := &proto.LockQueueMessage{}
				err = q.pdtool.Decode(b, QMsg)
				if err != nil {
					errtext = err.Error()
					finish <- true
					return
				}
				if QMsg.Id == q.QueueId && QMsg.Text == q.LockText {
					finish <- true
					return
				}

			}
			time.Sleep(time.Millisecond * time.Duration(timeSleep))
		}
	}()
	<-finish
	if errtext != "" {
		return errors.New("queue lock err: " + errtext)
	}
	//fmt.Println(q.GetKey(), q.GetId(), "加锁成功")

	return nil

}
func (q *Queue) UnLock() error {
	conn := q.IRedis.GetRawMasterConn()
	defer conn.Close()
	_, err := conn.Do("lrem", q.Key, 1, q.pushdata)
	if err != nil {
		return err
	}

	return nil
}
func (q *Queue) GetId() int64 {
	return q.QueueId
}
func (q *Queue) GetKey() string {
	return q.Key
}
func (q *Queue) GetTimeOut() int64 {
	return q.timeout
}
