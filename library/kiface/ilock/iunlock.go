package ilock

type IUnLock interface {
	/*
		解锁 强制解锁
		key string 所需要加锁的Key
		成功返回true
		失败返回false
	*/
	Do(Key string) bool
}
