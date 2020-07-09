package utils

import (
	"KServer/library/iface/utils"
	"encoding/json"
	"fmt"
)

type Json struct {
}

func NewIJson() utils.IJson {
	return &Json{}
}

// Struct转JSON
func (j *Json) StructToJson(tableStruct interface{}) (string, error) {
	b, err := json.Marshal(&tableStruct)
	if err != nil {
		fmt.Println("err", err)
		return "", err
	}
	return string(b), err
}

// Struct转BYTE
func (j *Json) StructToByte(tableStruct interface{}) ([]byte, error) {
	b, err := json.Marshal(&tableStruct)
	if err != nil {
		fmt.Println("err", err)
		return nil, err
	}
	return b, err
}

//JSON转Struct
func (j *Json) JsonToStruct(str string, tableStruct interface{}) error {

	if err := json.Unmarshal([]byte(str), &tableStruct); err == nil {
		return err
	} else {
		return err
	}

}
