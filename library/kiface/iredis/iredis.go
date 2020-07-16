package iredis

import "github.com/garyburd/redigo/redis"

type IRedisPool interface {
	StartMasterPool(addr string, password string, maxIdle int, maxActive int) bool
	StartSlavePool(addr string, password string, maxIdle int, maxActive int) bool
	GetMasterConn() IValue
	GetSlaveConn() IValue
	CloseMaster() error
	CloseSlave() error
	GetRawMasterConn() redis.Conn
	GetRawSlaveConn() redis.Conn
}

type IRedisConf interface {
	GetMasterAddr() string
	GetSlaveAddr() string
}
