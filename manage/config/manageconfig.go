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
	Socket struct {
		Client bool
		Server bool
	}
	// 是否开启client管理
}

func NewManageConfig() *ManageConfig {
	conf := &ManageConfig{}
	return conf
}
