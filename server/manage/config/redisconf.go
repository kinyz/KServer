package config

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

type RedisConfig struct {
	Env    bool   `yaml:"env"`
	Master Master `yaml:"master"`
	Slave  Slave  `yaml:"slave"`
}

type Master struct {
	Host      string `yaml:"host"`
	Port      string `yaml:"port"`
	PassWord  string `yaml:"password"`
	MaxIdle   int    `yaml:"maxIdle"`
	MaxActive int    `yaml:"maxActive"`
}

type Slave struct {
	Enable    bool   `yaml:"enable"`
	Host      string `yaml:"host"`
	Port      string `yaml:"port"`
	PassWord  string `yaml:"password"`
	MaxIdle   int    `yaml:"maxIdle"`
	MaxActive int    `yaml:"maxActive"`
}

// 取redis配置文件
func NewRedisConfig(fileName string) *RedisConfig {
	conf := &RedisConfig{}
	path, _ := os.Getwd()
	yamlFile, err := ioutil.ReadFile(path + fileName)

	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, &conf)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	if conf.Env {
		conf.Master.Host = os.Getenv("ENV_REDIS_MASTER_HOST")
		conf.Master.Port = os.Getenv("ENV_REDIS_MASTER_PORT")
		conf.Master.PassWord = os.Getenv("ENV_REDIS_MASTER_PASSWORD")
		conf.Master.MaxIdle, _ = strconv.Atoi(os.Getenv("ENV_REDIS_MASTER_MAXIDLE"))
		conf.Master.MaxActive, _ = strconv.Atoi(os.Getenv("ENV_REDIS_MASTER_MAXACTIVE"))

		conf.Slave.Host = os.Getenv("ENV_REDIS_SlAVE_HOST")
		conf.Slave.Port = os.Getenv("ENV_REDIS_SlAVE_PORT")
		conf.Slave.PassWord = os.Getenv("ENV_REDIS_SlAVE_PASSWORD")
		conf.Slave.MaxIdle, _ = strconv.Atoi(os.Getenv("ENV_REDIS_SlAVE_MAXIDLE"))
		conf.Slave.MaxActive, _ = strconv.Atoi(os.Getenv("ENV_REDIS_SlAVE_MAXACTIVE"))
	}
	return conf
}
func (c *RedisConfig) GetMasterAddr() string {
	return c.Master.Host + ":" + c.Master.Port
}
func (c *RedisConfig) GetSlaveAddr() string {
	return c.Master.Host + ":" + c.Master.Port
}
