package kafka

import (
	"KServer/library/iface/kafka"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"os"
)

type KafkaConf struct {
	Env  bool   `yaml:"env"`
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

func NewKafkaConf(filename string) kafka.IKafkaConf {
	conf := &KafkaConf{}
	path, _ := os.Getwd()
	yamlFile, err := ioutil.ReadFile(path + filename)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
		return nil
	}
	err = yaml.Unmarshal(yamlFile, &conf)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
		return nil
	}
	if conf.Env {
		conf.Host = os.Getenv("ENV_KAFKA_HOST")
		conf.Port = os.Getenv("ENV_KAFKA_PORT")
	}

	return conf
}

func (c *KafkaConf) GetAddr() string {
	return c.Host + ":" + c.Port
}
