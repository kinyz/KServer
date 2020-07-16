package ilock

type IQueue interface {
	/*
		获取队列锁
		timeSleep int 毫秒 设置每次获取队列的间隔时间
		timeOut int64 毫秒 设置无法获取队列退出时间
		成功返回nil
		失败返回error
	*/
	Lock(timeSleep int, timeOut int64) error
	/*
		队列解锁
		成功返回nil
		失败返回error
	*/
	UnLock() error
	/*
		获取队列ID
		返回值int64
	*/
	GetId() int64
	/*
		获取队列Key
		返回值string
	*/
	GetKey() string
	/*
		获取锁的时效
		返回值int64
	*/
	GetTimeOut() int64
}
