package utils

import (
	"KServer/library/iface/iutils"
	"encoding/json"
)

type Json struct {
}

func NewIJson() iutils.IJson {
	return &Json{}
}

func (j *Json) Encode(table interface{}) []byte {
	b, err := json.Marshal(&table)
	if err != nil {
		return nil
	}
	return b
}

func (j *Json) Decode(data []byte, table interface{}) error {

	return json.Unmarshal(data, &table)
}
