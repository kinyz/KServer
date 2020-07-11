package utils

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
)

type Yaml struct {
}

func (y *Yaml) ReadConfFileToStruct(fileName string, table interface{}) error {
	path, _ := os.Getwd()
	fmt.Println("配置文件地址=", path+fileName)
	yamlFile, err := ioutil.ReadFile(path + fileName)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(yamlFile, &table)
	if err != nil {
		return err
	}
	return nil
}
