package redis

type IRedisPool interface {
	StartMasterPool(addr string, password string, maxIdle int, maxActive int) bool
	StartSlavePool(addr string, password string, maxIdle int, maxActive int) bool
	GetMasterConn() IValue
	GetSlaveConn() IValue
	CloseMaster() error
	CloseSlave() error
}

type IRedisConf interface {
	GetMasterAddr() string
	GetSlaveAddr() string
}
