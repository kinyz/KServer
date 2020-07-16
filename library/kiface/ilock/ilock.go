package ilock

type ILock interface {

	/*
		锁 直锁
		key string 所需要加锁的Key
		成功返回true
		失败返回false
	*/
	Do(Key string) bool

	// 检查LockKey是否已存在
	Check(key string) bool

	/*
		自动锁
		key string 所需要加锁的Key
		timeOut int64 设置秒数后自动解锁
		成功返回true
		失败返回false
		自动解锁失败会返回日志
	*/
	Auto(key string, timeOut int) bool

	/*
		时间锁
		key string 所需要加锁的Key
		timeSleep int64 毫秒 设置毫秒进行一次加锁
		timeOut int64 毫秒 设置毫秒加锁返回结果
		成功返回true
		失败返回false
	*/
	Time(key string, timeSleep int, timeOut int) bool

	/*
		次数锁
		key string 所需要加锁的Key
		num int 设置次数内锁
		成功返回true
		失败返回false
	*/
	Num(key string, num int) bool

	/*
		队列锁
		key string 需要队列的key
		timeout int64 锁的时效时间
		返回一个队列对象
	*/
	Queue(key string, timeOut int64) IQueue
}
