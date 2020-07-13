package ilock

type ILock interface {
	Lock(Key string) bool
	UnLock(Key string) bool
}
