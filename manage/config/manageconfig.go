package config

import (
	"KServer/library/iface/ikafka"
	"KServer/server/utils"
)

type ManageConfig struct {
	// message开关
	Message struct {
		// 是否开启kafka管理
		Kafka bool
	}
	DB struct {
		// 是否开启redis管理
		Redis bool
		Mongo bool
	}
	Server struct {
		// 设置服务器头
		Head string
	}
	Socket struct {
		Client bool
		Server bool
	}
	WebSocket struct {
		Client bool
		Server bool
	}
	Lock struct {
		Open bool
		Head string
	}

	// 是否开启client管理
}

func NewManageConfig() *ManageConfig {
	conf := &ManageConfig{}

	//conf.Socket.Server =false
	return conf
}

func (m *ManageConfig) GetRedisConfig() *RedisConfig {
	return NewRedisConfig(utils.RedisConFile)
}
func (m *ManageConfig) GetKafkaConfig() ikafka.IKafkaConf {
	return NewKafkaConfig(utils.KafkaConFile)
}
