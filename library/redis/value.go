package redis

import (
	iface "KServer/library/iface/iredis"
	"github.com/garyburd/redigo/redis"
)

type Value struct {
	Conn redis.Conn
}

func (v *Value) Get(key string) iface.IGetValue {
	return &GetValue{Conn: v.Conn, Key: key}
}
func (v *Value) Set(key string) iface.ISetValue {
	return &SetValue{Conn: v.Conn, Key: key}
}
func (v *Value) Do(key string, value ...interface{}) (reply interface{}, err error) {
	return v.Conn.Do("SET", key, value)
}
func (v *Value) Check(key string) bool {

	defer v.Conn.Close()
	_, err := v.Conn.Do("EXISTS", key)
	if err != nil {
		return false
	}
	return true
}
