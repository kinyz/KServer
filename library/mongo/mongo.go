package mongo

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"os"
	"time"
)

var Session *mgo.Session
var MongoInterface Mongo

type Mongo interface {
	Init()
	GetSession() *mgo.Session
	GetCollection(collName string) *mgo.Collection
}

type IMongo struct {
	Env             bool   `yaml:"Env"`
	Name            string `yaml:"Name"`
	Host            string `yaml:"Host"`
	Port            string `yaml:"Port"`
	User            string `yaml:"User"`
	Password        string `yaml:"PassWord"`
	MaxReConnectNum int    `yaml:"MaxReConnectNum"`
}

func NewMongo() Mongo {
	MongoInterface = &IMongo{}
	return MongoInterface
}
func (m *IMongo) Init() {
	m.ReadConfFile("conf/mongo.yaml")
	if m.Env {
		m.Host = os.Getenv("ENV_MONGO_Host")
		m.Port = os.Getenv("ENV_MONGO_Port")
		m.Name = os.Getenv("ENV_MONGO_Name")
		m.User = os.Getenv("ENV_MONGO_User")
		m.Password = os.Getenv("ENV_MONGO_PassWord")
	}
	connectInfo := &mgo.DialInfo{
		Addrs:     []string{m.Host + ":" + m.Port},
		Direct:    false,
		Timeout:   time.Second * 1,
		Database:  m.Name,
		Source:    m.Name,
		Username:  m.User,
		Password:  m.Password,
		PoolLimit: 1024,
	}
	session, err := mgo.DialWithInfo(connectInfo)
	if err != nil {
		fmt.Println("MongoDB 数据库连接失败 失败信息:" + err.Error())
		//err_handler(err)
		return
	}
	Session = session
	Session.SetMode(mgo.Monotonic, true)
}

// 读取mongo配置文件
func (m *IMongo) ReadConfFile(fileName string) {
	yamlFile, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, m)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
}
func (m *IMongo) GetSession() *mgo.Session {
	err := Session.Ping()
	if err != nil {
		fmt.Println("MongoDB 数据库丢失连接 失败信息:" + err.Error())
		m.ReadConfFile("conf/mongo.yaml")
		if m.Env {
			m.Host = os.Getenv("ENV_MONGO_Host")
			m.Port = os.Getenv("ENV_MONGO_Port")
			m.Name = os.Getenv("ENV_MONGO_Name")
			m.User = os.Getenv("ENV_MONGO_User")
			m.Password = os.Getenv("ENV_MONGO_PassWord")
		}
		for i := 0; i <= m.MaxReConnectNum; i++ {
			fmt.Printf("MongoDB 正在进行重新连接...当前第 %d次", i)
			connectInfo := &mgo.DialInfo{
				Addrs:     []string{m.Host + ":" + m.Port},
				Direct:    false,
				Timeout:   time.Second * 1,
				Database:  m.Name,
				Source:    m.Name,
				Username:  m.User,
				Password:  m.Password,
				PoolLimit: 1024,
			}
			s, err := mgo.DialWithInfo(connectInfo)
			if err == nil {
				Session = s
				break
			}
			fmt.Println("[MongoDB] 连接失败 失败信息:" + err.Error())
			time.Sleep(1000)
			if i == m.MaxReConnectNum {
				fmt.Println("[MongoDB] 重连达到最大重连次数")
				return nil
			}
		}
	}
	return Session
}
func (m *IMongo) GetCollection(collName string) *mgo.Collection {
	m.ReadConfFile("conf/mongo.yaml")
	if m.Env {
		m.Name = os.Getenv("ENV_MONGO_Name")
	}
	s := m.GetSession()
	if s == nil {
		return nil
	}
	return s.DB(m.Name).C(collName)
}
