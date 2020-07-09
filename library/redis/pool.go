package redis

import (
	iface "KServer/library/iface/redis"
	"fmt"
	"github.com/garyburd/redigo/redis"
)

type Pool struct {
	MasterPool   *redis.Pool
	SlavePool    *redis.Pool
	MasterEnable bool
	SlaveEnable  bool
}

func NewIRedisPool() iface.IRedisPool {
	return &Pool{SlaveEnable: false, MasterEnable: false}
}

// 开启redis 主库
func (p *Pool) StartMasterPool(addr string, password string, maxIdle int, maxActive int) bool {
	if p.MasterEnable {
		fmt.Println("Redis MasterPool 已开启，请勿重复开启")
		return false
	}
	p.MasterPool = &redis.Pool{
		MaxIdle:     maxIdle,
		MaxActive:   maxActive,
		IdleTimeout: 100,
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial("tcp", addr, redis.DialPassword(password))
			if err != nil {
				fmt.Println("Redis 连接失败")
				//panic(err)
				return nil, err
			}
			return conn, nil
		},
	}
	p.MasterEnable = true
	return true
}

// 开启redis从库
func (p *Pool) StartSlavePool(addr string, password string, maxIdle int, maxActive int) bool {
	if !p.MasterEnable {
		fmt.Println("Redis MasterPool 未开启,请先开启")
		return false
	}
	if p.SlaveEnable {
		fmt.Println("Redis SlavePool 已开启，请勿重复开启")
		return false
	}
	p.SlavePool = &redis.Pool{
		MaxIdle:     maxIdle,
		MaxActive:   maxActive,
		IdleTimeout: 100,
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial("tcp", addr, redis.DialPassword(password))
			if err != nil {
				fmt.Println("Redis 连接失败", err)
				//panic(err)
				return nil, err
			}
			return conn, nil
		},
	}
	p.SlaveEnable = true
	return true
}
func (p *Pool) GetMasterPool() *redis.Pool {
	return p.MasterPool
}
func (p *Pool) GetSlavePool() *redis.Pool {
	return p.SlavePool
}

// 获取redis Master conn
func (p *Pool) GetMasterConn() iface.IValue {
	return &Value{Conn: p.GetMasterPool().Get()}
}

// 获取redis Slave conn
func (p *Pool) GetSlaveConn() iface.IValue {
	if !p.SlaveEnable {
		//	fmt.Println("我是master")
		return &Value{Conn: p.GetMasterPool().Get()}
	}
	//fmt.Println("我是slave")
	return &Value{Conn: p.GetSlavePool().Get()}
}

func (p *Pool) CloseMaster() error {
	//	p.GetMasterPool().Close()
	return p.GetMasterPool().Close()
}
func (p *Pool) CloseSlave() error {
	//p.GetMasterPool().Close()
	return p.GetSlavePool().Close()
}
