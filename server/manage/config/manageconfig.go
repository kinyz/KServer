package config

type ManageConfig struct {
	// message开关
	Message struct {
		// 是否开启kafka管理
		Kafka bool
	}
	DB struct {
		// 是否开启redis管理
		Redis bool
	}
	Server struct {
		// 设置服务器头
		Head string
	}
	// 是否开启client管理
	Client bool
}

func NewManageConfig() *ManageConfig {
	conf := &ManageConfig{}
	conf.Message.Kafka = false
	conf.DB.Redis = false
	conf.Client = false
	return conf
}
