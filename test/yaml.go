package main

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"os"
)

type Value struct {
	Topic map[string]struct {
		Key map[string]struct {
			Id map[uint32]uint32
		}
	}

	//Key   []string `yaml:"key"`
	//Id    []uint32 `yaml:"topic"`
}

type Msg struct {
}

func NewMsg() *Msg {
	return &Msg{}
}
func reload(fileName string) *Value {
	value := &Value{}
	path, _ := os.Getwd()
	fmt.Println(path + fileName)
	yamlFile, err := ioutil.ReadFile(path + fileName)

	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, &value)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return value
}

func main() {
	//v:=reload("/conf/msg.yaml")
	v := &Value{}

	v.Topic["1"].Key["2"].Id[2045] = 2

	//	fmt.Println(v.Topic)
	//m:=make(map[string]struct{Key map[string][]uint32})
	//m[""].Key["2"]=[]uint32{1,25,2,2,2,2}

	fmt.Println(v.Topic["1"].Key["2"].Id[2045])
}
